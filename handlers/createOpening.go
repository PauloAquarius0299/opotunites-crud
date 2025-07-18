package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mgs": "CreateOpeningHandler called",
	})
}