package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.MustGet("dst"))
}

//TODO: Написать хэндлер, который возвращает по айди растения его полный json
