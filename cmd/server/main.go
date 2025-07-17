package main

import (
	"db-service/internal/database"
	"db-service/internal/handler"
	"os"

	"log"

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
		port = "8000" // дефолтный порт если переменная окружения не задана
	}
	app := fiber.New()
	handler.RegisterTaskRoutes(app, pool)

	log.Printf("Server is running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
