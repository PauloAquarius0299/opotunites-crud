package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func handler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mgs": "handler routes called",
	})
}