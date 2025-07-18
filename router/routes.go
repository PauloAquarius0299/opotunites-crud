package router

import (
	"crud-oportunides/handlers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handlers.ListOpeningsHandler)
		v1.GET("/opening/:id", handlers.ShowOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.PUT("/opening/:id", handlers.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handlers.DeleteOpeningHandler)
	}
}