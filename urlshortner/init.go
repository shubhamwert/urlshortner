package main

import (
	"shubham/urlShortner/routes"

	"github.com/gofiber/fiber/v2"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	return app
}

func InitializeApp(app *fiber.App) {

	routes.InitRoutes(app)
}

func RunApp(app *fiber.App) {
	app.Listen(":8080")
}
