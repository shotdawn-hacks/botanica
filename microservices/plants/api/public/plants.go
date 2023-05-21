package public

import (
	"botanica/pkg/db"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "botanica"
	COLLECTION = "plants"
	IMAGES     = "images"
)

type Image struct {
	ID   string `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type PlantWithImage struct {
	Plant db.Plant
	Image string `json:"image" bson:"image"`
}

func Plants(ctx *gin.Context) {
	// firstname := ctx.DefaultQuery("firstname", "Guest")
	// lastname := ctx.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	mongoURI := ctx.GetString("MongoURI")
	if len(mongoURI) == 0 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, `{"reason":"mongo uri is not exsists"}`)
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, `{"reason":"mongo is not responding"}`)
		return
	}

	coll := client.Database(DATABASE).Collection(COLLECTION)
	imcoll := client.Database(DATABASE).Collection(IMAGES)

	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, `{"reason":"mongo error"}`)
		return
	}

	var allPlants []db.Plant

	var allWithImage []PlantWithImage

	err = cur.All(context.TODO(), &allPlants)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, `{"reason":"mongo error"}`)
		return
	}

	var img Image

	for _, plant := range allPlants {
		im := imcoll.FindOne(context.TODO(), bson.D{{"name", plant.Name}})
		err := im.Decode(&img)
		if err != nil {
			log.Println(err)
		}

		p := PlantWithImage{
			Plant: plant,
			Image: img.ID,
		}

		allWithImage = append(allWithImage, p)
	}

	ctx.JSON(http.StatusOK, allWithImage)
}

func PlantByUUID(ctx *gin.Context) {
	// Подключиться к монгоДБ
	// Выбрать коллекцию plants
	// Сделать поиск по uuid
	// При успехе сделать ctx.JSON(http.StatusOK, foundPlantJson)
	// При неудаче сделать ctx.AbortWithStatus(http.StatusBadRequest)
}

func UpdatePlantInformation(ctx *gin.Context) {

}
