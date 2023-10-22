package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prezeswastaken/election/handlers"
	"github.com/prezeswastaken/election/initializers"
)

func main() {
	initializers.ConnectToDB()
	initializers.Migrate()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/contestants", handlers.ListContestants)
	app.Post("/api/contestants/create", handlers.CreateContestant)
	app.Delete("/api/contestants/:id", handlers.DeleteContestant)

	app.Listen(":8000")
}
