package main

import (
	"api/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := os.Getenv("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
