package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const jsonchick = `{
    "_id": 15308,
    "usda": "ABAM",
    "scientific_name": "Abies amabilis",
    "common_name": "Тихоокеанская серебристая ель"
}`

func Plants(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsonchick)
}

func PlantByUUID(ctx *gin.Context) {
	// Подключиться к монгоДБ
	// Выбрать коллекцию plants
	// Сделать поиск по uuid
	// При успехе сделать ctx.JSON(http.StatusOK, foundPlantJson)
	// При неудаче сделать ctx.AbortWithStatus(http.StatusBadRequest)
}
