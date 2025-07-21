package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"error":     msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context,op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation %s was successful", op),
		"data":    data,
	})
	ctx.Status(http.StatusOK)
}