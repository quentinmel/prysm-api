package handlers

import "github.com/gofiber/fiber/v2"

// GetSwaggerJSON returns the OpenAPI 3.0 specification
func GetSwaggerJSON(c *fiber.Ctx) error {
    spec := fiber.Map{
        "openapi": "3.0.0",
        "info": fiber.Map{
            "title":       "Prysm API",
            "version":     "1.0.0",
            "description": "API REST - Plateforme de Prédictions Football",
        },
        "servers": []fiber.Map{
            {
                "url":         "http://localhost:3000/api/v1",
                "description": "Développement",
            },
        },
        "paths": fiber.Map{
            "/health": fiber.Map{
                "get": fiber.Map{
                    "tags":        []string{"System"},
                    "summary":     "Health check de l'API",
                    "description": "Vérifie l'état de l'API et les variables d'environnement",
                    "responses": fiber.Map{
                        "200": fiber.Map{
                            "description": "API opérationnelle",
                            "content": fiber.Map{
                                "application/json": fiber.Map{
                                    "schema": fiber.Map{
                                        "type": "object",
                                        "properties": fiber.Map{
                                            "success":   fiber.Map{"type": "boolean"},
                                            "timestamp": fiber.Map{"type": "string", "format": "date-time"},
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "/competitions": fiber.Map{
                "get": fiber.Map{
                    "tags":        []string{"Competitions"},
                    "summary":     "Liste des compétitions",
                    "description": "Retourne toutes les compétitions disponibles",
                    "parameters": []fiber.Map{
                        {
                            "name":        "status",
                            "in":          "query",
                            "description": "Filtrer par statut (active, inactive)",
                            "required":    false,
                            "schema":      fiber.Map{"type": "string"},
                        },
                        {
                            "name":        "country",
                            "in":          "query",
                            "description": "Filtrer par pays",
                            "required":    false,
                            "schema":      fiber.Map{"type": "string"},
                        },
                        {
                            "name":        "limit",
                            "in":          "query",
                            "description": "Nombre maximum de résultats",
                            "required":    false,
                            "schema":      fiber.Map{"type": "integer", "default": 50},
                        },
                    },
                    "responses": fiber.Map{
                        "200": fiber.Map{
                            "description": "Liste des compétitions récupérée avec succès",
                        },
                        "500": fiber.Map{
                            "description": "Erreur serveur",
                        },
                    },
                },
            },
            "/matches": fiber.Map{
                "get": fiber.Map{
                    "tags":        []string{"Matches"},
                    "summary":     "Liste des matchs",
                    "description": "Retourne tous les matchs avec filtres",
                    "parameters": []fiber.Map{
                        {
                            "name":        "status",
                            "in":          "query",
                            "description": "Filtrer par statut (open, closed, finished)",
                            "schema":      fiber.Map{"type": "string"},
                        },
                        {
                            "name":        "limit",
                            "in":          "query",
                            "description": "Nombre maximum de résultats",
                            "schema":      fiber.Map{"type": "integer", "default": 50},
                        },
                    },
                    "responses": fiber.Map{
                        "200": fiber.Map{
                            "description": "Liste des matchs",
                        },
                    },
                },
            },
            "/matches/{id}": fiber.Map{
                "get": fiber.Map{
                    "tags":        []string{"Matches"},
                    "summary":     "Détails d'un match",
                    "description": "Retourne les détails d'un match spécifique",
                    "parameters": []fiber.Map{
                        {
                            "name":        "id",
                            "in":          "path",
                            "description": "ID du match",
                            "required":    true,
                            "schema":      fiber.Map{"type": "string"},
                        },
                    },
                    "responses": fiber.Map{
                        "200": fiber.Map{
                            "description": "Détails du match",
                        },
                        "404": fiber.Map{
                            "description": "Match non trouvé",
                        },
                    },
                },
            },
        },
    }

    return c.JSON(spec)
}