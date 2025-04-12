package profiles

import (
	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/handler"
	"github.com/gin-gonic/gin"
)

type ProfileConfig interface {
	GetProfilesTablePath() string
}

func RegisterRoutes(cfg ProfileConfig, rg *gin.RouterGroup) {
	h := handler.New(cfg)
	profiles := rg.Group("/profiles")

	// List all profiles
	profiles.GET("", h.ListProfiles)

	// Get a specific profile
	profiles.GET("/:id", h.GetProfile)

	// Create a new profile
	profiles.POST("", h.CreateProfile)

	// Update a profile
	profiles.PUT("/:id", h.UpdateProfile)

	// Delete a profile
	profiles.DELETE("/:id", h.DeleteProfile)
}
