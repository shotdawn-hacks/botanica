package processor

import (
	"botanica/microservices/plants/api/public"
	"botanica/pkg/destination"

	"github.com/gin-gonic/gin"
)

type Plants struct {
	Address      string                     `json:"address" bson:"address"`
	Port         string                     `json:"port" bson:"port"`
	Destinations []*destination.Destination `json:"destinations" bson:"destinations"`
	Mongo        string
}

type Config struct {
	MongoURI string
}

func NewtDefaultPlants(cfgs Config) *Plants {
	c := &Plants{
		Address: "",
		Port:    "9000",
		Mongo:   cfgs.MongoURI,
	}

	return c
}

func (r *Plants) init() {
	for _, dest := range r.Destinations {
		dest.Start()
	}
}
func (r *Plants) SetMongoURI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("MongoURI", r.Mongo)
		ctx.Next()
	}
}

func (r *Plants) newAPI() *gin.Engine {
	router := gin.New()
	publicRouter := router.Group("/api/v1")

	//
	// STATIC
	//
	router.Static("/static", "./images")
	//
	// PLANTS
	//

	plants := publicRouter.Group("/plants", r.SetMongoURI())
	plants.GET(HTTPPlants, public.Plants)

	return router
}

func (r *Plants) Start() {
	r.init()

	api := r.newAPI()

	api.Run(":" + r.Port)
}
