package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mgs": "DeleteOpeningHandler called",
	})
}