package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/joho/godotenv"
    "github.com/quentinmel/prysm-api/handlers"
    "github.com/quentinmel/prysm-api/database"
)

func main() {
    // Charger les variables d'environnement
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialiser la base de données
    database.InitSupabase()

    // Créer l'application Fiber
    app := fiber.New(fiber.Config{
        AppName: "Prysm API v1.0.0",
    })

    // Middlewares
    app.Use(recover.New())
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,PUT,DELETE",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

    // Routes statiques pour Swagger UI
    app.Static("/swagger", "./static/swagger.html")

    // API routes
    api := app.Group("/api/v1")

    // Health check
    api.Get("/health", handlers.HealthCheck)

    // Competitions
    api.Get("/competitions", handlers.GetCompetitions)

    // Matches
    api.Get("/matches", handlers.GetMatches)
    api.Get("/matches/:id", handlers.GetMatchByID)

    // Swagger JSON spec
    api.Get("/swagger", handlers.GetSwaggerJSON)

    // Démarrer le serveur
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Printf("🚀 Server starting on port %s", port)
    log.Fatal(app.Listen(":" + port))
}