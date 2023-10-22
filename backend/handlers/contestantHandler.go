package handlers

import (
	"strconv"

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

func DeleteContestant(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	var contestant models.Contestant
	initializers.DB.First(&contestant, id)
	if contestant.ID == 0 {
		return c.Status(404).SendString("No contestant found with given ID")
	}

	initializers.DB.Delete(&contestant)
	return c.SendString("Contestant successfully deleted")
}
