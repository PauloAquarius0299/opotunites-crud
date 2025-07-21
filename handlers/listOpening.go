package handlers

import (
	"crud-oportunides/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error retrieving openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}