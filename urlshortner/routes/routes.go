package routes

import (
	"fmt"
	"net/http"
	"shubham/urlShortner/controller"

	"github.com/gofiber/fiber/v2"
)

var handler *controller.UrlController

func InitRoutes(app *fiber.App, H *controller.UrlController) {
	handler = H
	app.Get("/", getIndex)
	app.Get("/new/:url/:customUrl?", postCreateShortUrl)
	app.Get("/get/:url", GetShortenUrl)

}

func getIndex(c *fiber.Ctx) error {
	c.SendString("Lets Get Started")
	return nil
}

func postCreateShortUrl(c *fiber.Ctx) error {
	fmt.Println(c.Get("x-auth-user"))
	s, err := handler.Shorten(c.Context(), c.Params("url"), c.Get("x-auth-user", "test"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	c.SendString(s)
	return nil
}
func GetShortenUrl(c *fiber.Ctx) error {
	fmt.Println()

	p := c.Params("url")
	u, err := handler.GetUrl(c.Context(), p, c.Get("x-auth-user", "test"))
	if u == "" {
		c.SendStatus(404)
		return nil
	}
	if err != nil {
		fmt.Println(err)

	}
	err = c.Redirect(
		fmt.Sprintf("https://www.%s", u), http.StatusPermanentRedirect,
	)
	if err != nil {
		fmt.Println(err)

	}
	return nil
}
