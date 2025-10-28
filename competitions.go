package handlers

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/quentinmel/prysm-api/database"
    "github.com/quentinmel/prysm-api/models"
)

func GetCompetitions(c *fiber.Ctx) error {
    status := c.Query("status")
    country := c.Query("country", "International")
    limitStr := c.Query("limit", "50")
    
    limit, _ := strconv.Atoi(limitStr)

    var rooms []models.Room
    query := database.Client.From("rooms").
        Select("*", "exact", false).
        Order("match_date", &map[string]interface{}{"ascending": false}).
        Limit(uint(limit), "")

    if status != "" {
        query = query.Eq("status", status)
    }

    _, count, err := query.Execute(&rooms)
    if err != nil {
        return c.Status(500).JSON(models.ErrorResponse{
            Success: false,
            Message: "Erreur lors de la récupération des compétitions",
        })
    }

    // Grouper par compétition
    competitionsMap := make(map[string]*models.Competition)
    
    for _, room := range rooms {
        key := room.TeamHome + "-" + room.TeamAway
        
        if _, exists := competitionsMap[key]; !exists {
            competitionsMap[key] = &models.Competition{
                ID:                room.ID,
                Name:              room.TeamHome + " vs " + room.TeamAway,
                Type:              "football",
                Country:           country,
                Status:            room.Status,
                Teams:             []string{room.TeamHome, room.TeamAway},
                MatchDate:         room.MatchDate,
                CreatedAt:         room.CreatedAt,
                TotalRooms:        1,
                TotalParticipants: 0,
            }
        } else {
            competitionsMap[key].TotalRooms++
        }
    }

    competitions := make([]models.Competition, 0, len(competitionsMap))
    for _, comp := range competitionsMap {
        competitions = append(competitions, *comp)
    }

    return c.JSON(models.SuccessResponse{
        Success: true,
        Data:    competitions,
        Meta: &models.Meta{
            Total: int(count),
            Page:  1,
            Limit: limit,
        },
    })
}