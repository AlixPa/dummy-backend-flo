package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/service"
	"github.com/AlixPa/dummy-backend-flo/internal/common"
)

type Handler struct {
	s *service.Service
}

type HandlerConfig interface {
	GetDbTablesCsvPath() common.DbTablesCsv
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{s: service.New(cfg)}
}

func (h Handler) ListProfiles(c *gin.Context) {
	response, err := h.s.ListProfiles()
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

	if err := h.s.CreateProfile(req.Name, req.Age); err != nil {
		if errors.Is(err, common.ErrDuplicateProfileName) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile created successfully"})
}
