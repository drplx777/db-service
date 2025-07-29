package main

import (
	"db-service/internal/database"
	"db-service/internal/handler"
	"db-service/internal/handler/middleware"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	pool, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}
	defer database.CloseDB(pool)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := fiber.New()
	app.Use(middleware.SlogLogger())
	handler.RegisterTaskRoutes(app, pool)
	handler.RegisterUserRoutes(app, pool)

	slog.Info("Service started", "port", 8080)
	slog.Warn("Low disk space", "disk", "/dev/sda1", "free_percent", 5)
	slog.Error("Database connection failed", "error", "timeout")

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
