// users/routes.go
package ping

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/ping/handler"
	"github.com/gin-gonic/gin" // or echo, chi, etc.
)

func RegisterRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")
	ping.GET("", handler.Ping)
}
