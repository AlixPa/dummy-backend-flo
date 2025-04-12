package handler

import (
	"net/http"

	"github.com/AlixPa/dummy-backend-flo/internal/api/ping/service"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response := service.Pong()
	c.JSON(http.StatusOK, gin.H{"message": response})
}
