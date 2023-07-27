package main

import (
	"album-list/database"
	"album-list/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// go work use .
	database.ConnectDb()

	// client := database.DB.Db

	// coll := client.Database("album-list").Collection("albums")
	// // var result bson.M
	// result, err := coll.Find(context.TODO(), bson.D{})

	// var results []bson.M

	// if err = result.All(context.TODO(), &results); err != nil {
	// 	panic(err)
	// }

	// for _, result := range results {
	// 	fmt.Println(result)
	// }

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
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("stan Weeekly for clear skin or else nerd!")
	// })

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
