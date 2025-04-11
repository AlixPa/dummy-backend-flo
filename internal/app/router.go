package app

import (
	"path"

	"github.com/AlixPa/dummy-backend-flo/internal/api/ping"
	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles"
	"github.com/AlixPa/dummy-backend-flo/internal/common"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *common.Config) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")
	ping.RegisterRoutes(api)
	profiles.RegisterRoutes(profiles.Config{DbCsvPath: path.Join(cfg.DbPath, "profiles.csv")}, api)

	return router
}
