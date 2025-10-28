package handlers

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/quentinmel/prysm-api/database"
    "github.com/quentinmel/prysm-api/models"
)

func GetMatches(c *fiber.Ctx) error {
    status := c.Query("status")
    team := c.Query("team")
    dateFrom := c.Query("date_from")
    dateTo := c.Query("date_to")
    limitStr := c.Query("limit", "50")
    pageStr := c.Query("page", "1")

    limit, _ := strconv.Atoi(limitStr)
    page, _ := strconv.Atoi(pageStr)
    offset := (page - 1) * limit

    var rooms []models.Room
    query := database.Client.From("rooms").
        Select("*", "exact", false).
        Order("match_date", &map[string]interface{}{"ascending": false}).
        Range(uint(offset), uint(offset+limit-1), "")

    if status != "" {
        query = query.Eq("status", status)
    }

    if team != "" {
        query = query.Or("team_home.ilike.%"+team+"%,team_away.ilike.%"+team+"%", "")
    }

    if dateFrom != "" {
        query = query.Gte("match_date", dateFrom)
    }

    if dateTo != "" {
        query = query.Lte("match_date", dateTo)
    }

    _, count, err := query.Execute(&rooms)
    if err != nil {
        return c.Status(500).JSON(models.ErrorResponse{
            Success: false,
            Message: "Erreur lors de la récupération des matchs",
        })
    }

    totalPages := int(count) / limit
    if int(count)%limit != 0 {
        totalPages++
    }

    return c.JSON(models.SuccessResponse{
        Success: true,
        Data:    rooms,
        Meta: &models.Meta{
            Total:      int(count),
            Page:       page,
            Limit:      limit,
            TotalPages: totalPages,
        },
    })
}

func GetMatchByID(c *fiber.Ctx) error {
    matchID := c.Params("id")
    includePredictions := c.Query("include_predictions", "false")

    var room models.Room
    _, _, err := database.Client.From("rooms").
        Select("*", "", false).
        Eq("id", matchID).
        Single().
        Execute(&room)

    if err != nil {
        return c.Status(404).JSON(models.ErrorResponse{
            Success: false,
            Message: "Match non trouvé",
        })
    }

    match := models.Match{
        Room:             room,
        PredictionsCount: 0,
        Predictions:      []models.Prediction{},
    }

    if includePredictions == "true" {
        // TODO: Récupérer les prédictions depuis la table predictions
        // Pour l'instant, on laisse vide
    }

    return c.JSON(models.SuccessResponse{
        Success: true,
        Data:    match,
    })
}