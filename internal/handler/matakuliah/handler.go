package matakuliah

import (
	"api/internal/model"
	"api/internal/repository/db"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataMatakuliah(c *fiber.Ctx) error {
	var requestData model.Matakuliah
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	requestData.ID = uuid.New().String()

	// If data doesn't exist, proceed with insertion
	if err := db.InsertDataMatakuliah(requestData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data into the database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
		"data":    requestData,
	})
}

func GetAllDataMatakuliah(c *fiber.Ctx) error {
	filter := bson.M{}

	requestData, err := db.GetDataMatakuliah(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}
