package main

import (
	"cv-service/config"
	"cv-service/internal/grpc"
	"cv-service/pkg/logger"
	middleware "cv-service/rest/middlewares"
	"cv-service/rest/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	grpcConfig := &grpc.Config{
		AuthServiceAddress: "auth-service:50051",
	}

	clients, err := grpc.NewClients(grpcConfig)
	if err != nil {
		log.Fatal("Failed to create grpc clients:", err)
	}
	defer clients.Auth.Close()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	log := config.InitLogger(cfg)
	log.Debug("Configuration loaded : ", cfg)

	db, err := config.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer config.CloseDB(db)

	recommendationHistoryController := InitializeRecommendationHistoryDependency(db, clients.Auth)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		AppName:      "CV Service",
	})

	app.Use(cors.New())
	app.Use(logger.LogrusMiddleware(log))

	app.Get("/health", func(c *fiber.Ctx) error {
		log.Info("Health check endpoint called")
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Healthy",
		})
	})

	routes.RegisterRecommendationHistoryRoutes(app, recommendationHistoryController, clients.Auth)

	log.Info(fmt.Sprintf("Starting server on port %s", cfg.ServerPort))

	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
