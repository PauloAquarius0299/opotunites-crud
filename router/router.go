package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	// Roda o servidor
	router.Run(":8080")
}
