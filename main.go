package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main", // add this to config
	})

	setupRoutes(app)

	app.Listen(":3000")

}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("stan Weeekly for clear skin!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(http.StatusOK)
	})
}
