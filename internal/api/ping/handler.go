package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response := GetPingResponse()
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
