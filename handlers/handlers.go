package handlers

import (
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/quentinmel/prysm-api/database"
    "github.com/quentinmel/prysm-api/models"
    "github.com/supabase-community/postgrest-go"
)

type HealthCheckResponse struct {
    Success   bool      `json:"success"`
    Timestamp time.Time `json:"timestamp"`
    EnvCheck  EnvCheck  `json:"env_check"`
}

type EnvCheck struct {
    SupabaseURLExists   bool   `json:"supabase_url_exists"`
    SupabaseURLValue    string `json:"supabase_url_value"`
    SupabaseKeyExists   bool   `json:"supabase_key_exists"`
    SupabaseKeyLength   int    `json:"supabase_key_length"`
}

func HealthCheck(c *fiber.Ctx) error {
    supabaseURL := os.Getenv("SUPABASE_URL")
    supabaseKey := os.Getenv("SUPABASE_KEY")

    return c.JSON(HealthCheckResponse{
        Success:   true,
        Timestamp: time.Now(),
        EnvCheck: EnvCheck{
            SupabaseURLExists: supabaseURL != "",
            SupabaseURLValue:  supabaseURL,
            SupabaseKeyExists: supabaseKey != "",
            SupabaseKeyLength: len(supabaseKey),
        },
    })
}

// GetCompetitions returns all competitions
func GetCompetitions(c *fiber.Ctx) error {
    // Récupérer les paramètres de query
    status := c.Query("status")
    country := c.Query("country", "International")
    limit := c.QueryInt("limit", 50)

    // Query de base
    query := database.Client.From("rooms").
        Select("*", "exact", false).
        Order("match_date", &postgrest.OrderOpts{Ascending: false}).
        Limit(limit, "")

    // Appliquer les filtres
    if status != "" {
        query = query.Eq("status", status)
    }

    var rooms []models.Room
    count, err := query.ExecuteTo(&rooms)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "error":   err.Error(),
        })
    }

    // Grouper par compétition (team_home-team_away)
    competitionsMap := make(map[string]*models.Competition)

    for _, room := range rooms {
        key := room.TeamHome + "-" + room.TeamAway

        if comp, exists := competitionsMap[key]; exists {
            comp.TotalRooms++
        } else {
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
        }
    }

    // Convertir map en slice
    competitions := make([]models.Competition, 0, len(competitionsMap))
    for _, comp := range competitionsMap {
        competitions = append(competitions, *comp)
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    competitions,
        "meta": fiber.Map{
            "total": count,
            "page":  1,
            "limit": limit,
        },
    })
}

// GetMatches returns all matches
func GetMatches(c *fiber.Ctx) error {
    status := c.Query("status")
    limit := c.QueryInt("limit", 50)

    query := database.Client.From("rooms").
        Select("*", "exact", false).
        Order("match_date", &postgrest.OrderOpts{Ascending: false}).
        Limit(limit, "")

    if status != "" {
        query = query.Eq("status", status)
    }

    var matches []models.Match
    count, err := query.ExecuteTo(&matches)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "error":   err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    matches,
        "meta": fiber.Map{
            "total": count,
            "page":  1,
            "limit": limit,
        },
    })
}

// GetMatchByID returns a specific match by ID
func GetMatchByID(c *fiber.Ctx) error {
    id := c.Params("id")

    var match models.Match
    _, err := database.Client.From("rooms").
        Select("*", "", false).
        Eq("id", id).
        Single().
        ExecuteTo(&match)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "success": false,
            "error":   "Match not found",
        })
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    match,
    })
}