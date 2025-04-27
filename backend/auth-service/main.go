package main

import (
	"auth-service/config"
	"auth-service/internal/grpc"
	"auth-service/pkg/logger"
	middleware "auth-service/rest/middlewares"
	"auth-service/rest/routes"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Setup gRPC server
	grpcServer := grpc.NewServer(50051)
	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Fatal("Failed to start gRPC server:", err)
		}
	}()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	log := config.InitLogger(cfg)
	log.Debug("Configuration loaded : ", cfg)

	// Initialize google oauth
	googleAuthCfg, err := config.InitGoogleAuth(cfg)
	if err != nil {
		log.Fatal("Failed to initialize Google OAuth:", err)
	}

	db, err := config.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer config.CloseDB(db)

	authController := InitializeUserDependency(db, googleAuthCfg)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		AppName:      "Auth Service",
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

	routes.RegisterUserRoutes(app, authController)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info("Shutting down servers...")
		grpcServer.Stop()
		if err := app.Shutdown(); err != nil {
			log.Error("Error during shutdown:", err)
		}
	}()

	log.Info(fmt.Sprintf("Starting server on port %s", cfg.ServerPort))

	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}
}
