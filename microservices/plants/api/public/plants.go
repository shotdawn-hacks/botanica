package public

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Plant struct {
	ID               string   `json:"_id"`
	Name             string   `json:"name"`
	ScientificName   string   `json:"scientific_name"`
	Family           string   `json:"family"`
	ScientificFamily string   `json:"scientific_family"`
	Habitats         []string `json:"habitats"`
	Regions          []any    `json:"regions"`
	Soils            []string `json:"soils"`
	IsGlobalRedbook  bool     `json:"is_global_redbook"`
	Sowing           struct {
		Start       string `json:"start"`
		End         string `json:"end"`
		Conditions  string `json:"conditions"`
		Description string `json:"description"`
	} `json:"sowing"`
	Harvesting struct {
		Start       string `json:"start"`
		End         string `json:"end"`
		Conditions  string `json:"conditions"`
		Description string `json:"description"`
	} `json:"harvesting"`
	ChemicalComposition []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"chemical_composition"`
	MedicalPreparations string `json:"medical_preparations"`
	AnnualDemandTons    int    `json:"annual_demand_tons"`
}

const (
	DATABASE   = "botanica"
	COLLECTION = "plants"
)

func Plants(ctx *gin.Context) {
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

	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, `{"reason":"mongo error"}`)
		return
	}

	var allPlants []Plant

	err = cur.All(context.TODO(), &allPlants)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, `{"reason":"mongo error"}`)
		return
	}

	ctx.JSON(http.StatusOK, allPlants)
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
