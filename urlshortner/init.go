package main

import (
	"shubham/urlShortner/controller"
	"shubham/urlShortner/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateApp() *fiber.App {
	app := fiber.New()
	return app
}

func InitializeApp(app *fiber.App, handler *controller.UrlController) {

	routes.InitRoutes(app, handler)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}

func RunApp(app *fiber.App) {
	app.Listen(":9080")
}
