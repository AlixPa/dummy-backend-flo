package handler

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/ping/service"
	"github.com/AlixPa/dummy-backend-flo/internal/common/response"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	pong := service.Pong()
	response.SendOK(c, response.Message{Message: pong})
}
