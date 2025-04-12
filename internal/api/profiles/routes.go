package profiles

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/handler"
	"github.com/AlixPa/dummy-backend-flo/internal/common"
	"github.com/gin-gonic/gin"
)

type ProfileConfig interface {
	GetDbTablesCsvPath() common.DbTablesCsv
}

func RegisterRoutes(cfg ProfileConfig, rg *gin.RouterGroup) {
	h := handler.New(cfg)
	profiles := rg.Group("/profiles")

	profiles.GET("", h.ListProfiles)
	profiles.POST("", h.CreateProfile)
}
