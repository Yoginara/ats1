package router

import (
	"api/internal/handler/matakuliah" // Mengganti matakuliah menjadi matakuliah
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/matakuliah", matakuliah.InsertDataMatakuliah) // Mengubah menjadi matakuliah
	api.Get("/matakuliah", matakuliah.GetAllDataMatakuliah)  // Mengubah menjadi matakuliah
}
