package main

import (
	"album-list/database"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	database.ConnectDb()

	// client := database.DB.Db

	// coll := client.Database("album-list").Collection("albums")
	// name := "We Play"
	// var result bson.M
	// err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the name %s\n", name)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

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
		return c.SendString("stan Weeekly for clear skin or else!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(http.StatusOK)
	})
}
