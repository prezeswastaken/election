package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prezeswastaken/election/initializers"
	"github.com/prezeswastaken/election/models"
)

func CreateContestant(c *fiber.Ctx) error {
	contestant := new(models.Contestant)
	if err := c.BodyParser(contestant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	result := initializers.DB.Create(&contestant)
	if result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.Status(200).JSON(contestant)
}

func ListContestants(c *fiber.Ctx) error {
	contestants := []models.Contestant{}
	if result := initializers.DB.Find(&contestants); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(contestants)
}
