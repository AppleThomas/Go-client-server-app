package main

import (
	"album-list/database"
	"album-list/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main", // add this to config
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")

}

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListAlbums)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(http.StatusOK)
	})

	app.Get("/album", handlers.NewAlbumView)
	app.Post("/album", handlers.AddAlbum)

	app.Get("/album/:id", handlers.ShowAlbum)

	app.Get("/album/:id/edit", handlers.EditAlbum)
	app.Patch("/album/:id", handlers.UpdateAlbum)

	app.Delete("/album/:id", handlers.DeleteAlbum)
}
