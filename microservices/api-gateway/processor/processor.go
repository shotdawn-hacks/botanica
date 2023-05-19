package processor

import (
	"botanica/microservices/api-gateway/api/private"
	"botanica/microservices/api-gateway/api/public"
	"botanica/pkg/destination"
	"botanica/pkg/shared"

	"github.com/gin-gonic/gin"
)

type Gateway struct {
	ID           string                    `json:"_id" bson:"_id"`
	Address      string                    `json:"address" bson:"address"`
	Port         string                    `json:"port" bson:"port"`
	Destinations []destination.Destination `json:"destinations" bson:"destinations"`
}

type Config struct {
	Plants *destination.Config
}

func NewtDefaultGateway(cfgs Config) *Gateway {
	c := &Gateway{
		Address: "",
		Port:    "9000",
	}

	return c
}

func (r *Gateway) init() {
	for _, dest := range r.Destinations {
		err := dest.Client.Start()
		if err != nil {
			panic(err)
		}
	}
}

func (r *Gateway) Start() {
	r.init()

	api := r.newAPI()

	api.Run(":" + r.Port)
}

func (r *Gateway) SetPlantsDestiantion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("dst", r)
		ctx.Next()
	}
}

func (r *Gateway) newAPI() *gin.Engine {
	router := gin.New()
	publicRouter := router.Group("/api/v1")

	//
	// CORS
	//

	router.Use(shared.SetDefaultCors())

	//
	// PLANTS
	//

	plants := publicRouter.Group("/plants", r.SetPlantsDestiantion())
	plants.GET(HTTPPlantSearch, public.Search)

	//
	// PRIVATE
	//

	router.GET(HTTPHealth, private.Health)
	router.GET("/ping", private.Ping)

	return router
}
