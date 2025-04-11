// users/routes.go
package profiles

import (
	h "github.com/AlixPa/dummy-backend-flo/internal/api/profiles/handler"
	s "github.com/AlixPa/dummy-backend-flo/internal/api/profiles/service"
	"github.com/gin-gonic/gin"
)

type Config = s.Config

func RegisterRoutes(cfg Config, rg *gin.RouterGroup) {
	handler := h.New(cfg)
	profiles := rg.Group("/profiles")

	profiles.GET("/", handler.ListProfiles)
}
