package server

import (
	"botanica/microservices/plants/api/public"
	"botanica/pkg/shared"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	//
	// CORS
	//

	router.Use(shared.SetDefaultCors())

	//
	// PLANTS
	//

	router.GET(HTTPPlants, public.Plants)

	return router
}
