package shared

import (
	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ReadBodyAndDecode(ctx *gin.Context, v interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(&v)
}

func MustReadBodyAndDecode(ctx *gin.Context, v interface{}) {
	if err := ReadBodyAndDecode(ctx, v); err != nil {
		panic(err)
	}
}

func SetDefaultCors() gin.HandlerFunc {

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	return cors.New(corsConfig)
}

func SetCustomCors(cfg cors.Config) gin.HandlerFunc {
	return cors.New(cfg)
}
