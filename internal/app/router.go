package app

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/ping"
	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles"
	"github.com/gin-gonic/gin"
)

type routeurConfig interface {
	GetProfilesTablePath() string
}

func SetupRouter(cfg routeurConfig) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")
	ping.RegisterRoutes(api)
	profiles.RegisterRoutes(cfg, api)

	return router
}
