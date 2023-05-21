package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	_, ok := ctx.Get("dst")
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// iDst := dst.(processor.Gateway)
}
