package handlers

import (
	"album-list/database"
	"album-list/models"
	"context"
	"strings"

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
		return NewAlbumView(c)
	}

	data := bson.D{{"name", album.Name}, {"group", album.Group}, {"year", album.Year}, {"songs", album.Songs}, {"img", album.Img}}

	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}

	return ListAlbums(c)
}

func NewAlbumView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Album",
		"Subtitle": "Add your nifty new album from your collection",
	})
}

func ShowAlbum(c *fiber.Ctx) error {
	// album := models.Album{}
	id := c.Params("id")
	// id_url := c.OriginalURL()
	// fmt.Println("first")
	// fmt.Println(id_url)
	// id := id_url[len(id_url)-24:]
	// fmt.Println(id)

	objectId, err := primitive.ObjectIDFromHex(id)

	// fmt.Println(objectId)
	// fmt.Println(c.OriginalURL())

	// if err != nil {
	// 	panic(err)
	// }

	client := database.DB.Db
	var album models.Album

	coll := client.Database("album-list").Collection("albums")
	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&album)
	if err != nil {
		return NotFound(c)
	}

	songs := album.Songs
	multi := strings.Split(songs, ",")
	// fmt.Println(reflect.TypeOf(multi))
	album.SongsMulti = multi

	return c.Render("show", fiber.Map{
		"Title": album.Name,
		"Album": album,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")

}
