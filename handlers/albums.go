package handlers

import (
	"album-list/database"
	"album-list/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListAlbums(c *fiber.Ctx) error {

	client := database.DB.Db

	coll := client.Database("album-list").Collection("albums")
	result, err := coll.Find(context.TODO(), bson.D{})

	var albums []bson.M

	if err = result.All(context.TODO(), &albums); err != nil {
		panic(err)
	}

	var albums_ = make([]models.Album, len(albums))

	for index, result := range albums {
		albums_[index].ID = result["_id"].(primitive.ObjectID).Hex()
		albums_[index].Name = result["name"].(string)
		albums_[index].Group = result["group"].(string)
		albums_[index].Songs = result["songs"].(string)
		albums_[index].Img = result["img"].(string)
		albums_[index].Year = result["year"].(string)
	}

	return c.Render("index", fiber.Map{
		"Title":    "Kpop Album List",
		"Subtitle": "A place for you to keep track of all your kpop albums of the groups you stan",
		"Albums":   albums_,
	})
}

func AddAlbum(c *fiber.Ctx) error {

	client := database.DB.Db

	coll := client.Database("album-list").Collection("albums")

	album := new(models.Album)

	if err := c.BodyParser(album); err != nil {
		// fmt.Println("Start here 1")
		return NewAlbumView(c)
	}

	data := bson.D{{"name", album.Name}, {"group", album.Group}, {"year", album.Year}, {"songs", album.Songs}, {"img", album.Img}}

	_, err := coll.InsertOne(context.TODO(), data)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}

	// fmt.Println(album.Name)
	// fmt.Println(album.Group)
	// fmt.Println(album.Year)
	// fmt.Println(album.Songs)
	// fmt.Println(album.Img)

	return ListAlbums(c)
}

func NewAlbumView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Album",
		"Subtitle": "Add your nifty new album from your collection",
	})
}
