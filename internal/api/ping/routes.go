// users/routes.go
package ping

import (
	"github.com/gin-gonic/gin" // or echo, chi, etc.
)

func RegisterRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")
	ping.GET("/", Ping)
}
