package handler

import (
	"net/http"

	s "github.com/AlixPa/dummy-backend-flo/internal/api/ping/service"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response := s.Pong()
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
