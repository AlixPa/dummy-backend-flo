package app

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/ping"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")
	ping.RegisterRoutes(api)

	return router
}
