package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	app.Get("/", getIndex)
}

func getIndex(c *fiber.Ctx) error {
	c.SendString("Lets Get Started")
	return nil
}
