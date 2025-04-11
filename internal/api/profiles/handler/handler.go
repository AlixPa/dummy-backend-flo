package handler

import (
	"errors"
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
	response, err := h.service.ListProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
	}
}

func (h Handler) CreateProfile(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProfile(req.Name, req.Age); err != nil {
		if errors.Is(err, s.ErrDuplicateProfileName) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.Status(http.StatusOK)
}
