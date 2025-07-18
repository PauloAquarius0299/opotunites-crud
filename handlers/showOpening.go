package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mgs": "ShowOpeningHandler called",
	})
}
