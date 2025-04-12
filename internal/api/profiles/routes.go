// users/routes.go
package profiles

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/handler"
	"github.com/AlixPa/dummy-backend-flo/internal/common"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(cfg *common.Config, rg *gin.RouterGroup) {
	h := handler.New(cfg)
	profiles := rg.Group("/profiles")

	profiles.GET("/", h.ListProfiles)
	profiles.POST("/", h.CreateProfile)
}
