package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	s "github.com/AlixPa/dummy-backend-flo/internal/api/profiles/service"
)

type Handler struct {
	service *s.Service
}

func New(cfg s.Config) *Handler {
	return &Handler{service: s.New(cfg)}
}

func (h Handler) ListProfiles(c *gin.Context) {
	response := h.service.ListProfiles()
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
