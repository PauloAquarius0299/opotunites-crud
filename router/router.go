package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa a router do Gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
	// This will start the server on port 8080
}