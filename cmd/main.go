package main

import (
	"github.com/antpard/fiber-pscale/handlers"
	"github.com/antpard/fiber-pscale/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Fiber with Planetscale",
		ServerHeader: "Fiber",
	})
	models.ConnectDatabase()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Hello World",
		})
	})
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)
	app.Post("/users", handlers.CreateUser)
	app.Listen(":3000")
}
