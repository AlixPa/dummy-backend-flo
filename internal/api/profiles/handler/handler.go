package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/model"
	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/service"
	"github.com/AlixPa/dummy-backend-flo/internal/common/response"
)

type Handler struct {
	s *service.Service
}

type HandlerConfig interface {
	GetProfilesTablePath() string
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{s: service.New(cfg)}
}

// ListProfiles handles GET /api/profiles
func (h Handler) ListProfiles(c *gin.Context) {
	profiles, err := h.s.ListProfiles()
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to list profiles", err.Error())
		return
	}
	response.SendOK(c, response.Data{Data: profiles})
}

// GetProfile handles GET /api/profiles/:id
func (h Handler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.s.GetProfile(id)
	if err != nil {
		if errors.Is(err, service.ErrProfileNotFound) {
			response.SendError(c, http.StatusNotFound, "Profile not found", err.Error())
		} else {
			response.SendError(c, http.StatusInternalServerError, "Failed to get profile", err.Error())
		}
		return
	}
	response.SendOK(c, response.Data{Data: profile})
}

// CreateProfile handles POST /api/profiles
func (h Handler) CreateProfile(c *gin.Context) {
	var profile model.ProfileFields

	if err := c.ShouldBindJSON(&profile); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.s.CreateProfile(profile); err != nil {
		var validationErr *model.ValidationError
		if errors.As(err, &validationErr) {
			response.SendError(c, http.StatusBadRequest, "Invalid profile data", validationErr.Error())
		} else if errors.Is(err, service.ErrDuplicateProfileName) {
			response.SendError(c, http.StatusBadRequest, "Profile already exists", err.Error())
		} else {
			response.SendError(c, http.StatusInternalServerError, "Failed to create profile", err.Error())
		}
		return
	}
	response.SendCreated(c, response.Message{Message: "Profile created successfully"})
}

// UpdateProfile handles PUT /api/profiles/:id
func (h Handler) UpdateProfile(c *gin.Context) {
	id := c.Param("id")

	var profile model.ProfileFields

	if err := c.ShouldBindJSON(&profile); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.s.UpdateProfile(id, profile); err != nil {
		var validationErr *model.ValidationError
		if errors.As(err, &validationErr) {
			response.SendError(c, http.StatusBadRequest, "Invalid profile data", validationErr.Error())
		} else if errors.Is(err, service.ErrProfileNotFound) {
			response.SendError(c, http.StatusNotFound, "Profile not found", err.Error())
		} else {
			response.SendError(c, http.StatusInternalServerError, "Failed to update profile", err.Error())
		}
		return
	}
	response.SendNoContent(c)
}

// DeleteProfile handles DELETE /api/profiles/:id
func (h Handler) DeleteProfile(c *gin.Context) {
	id := c.Param("id")

	if err := h.s.DeleteProfile(id); err != nil {
		if errors.Is(err, service.ErrProfileNotFound) {
			response.SendError(c, http.StatusNotFound, "Profile not found", err.Error())
		} else {
			response.SendError(c, http.StatusInternalServerError, "Failed to delete profile", err.Error())
		}
		return
	}
	response.SendNoContent(c)
}
