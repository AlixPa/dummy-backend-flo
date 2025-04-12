package app

import (
	"log"

	"github.com/AlixPa/dummy-backend-flo/internal/common"
	"github.com/gin-gonic/gin"
)

func Run() {
	cfg, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Set Gin mode from configuration
	gin.SetMode(cfg.GetGinMode())

	router := SetupRouter(cfg)
	log.Printf("Server starting on port %s", cfg.GetPort())
	if err := router.Run(":" + cfg.GetPort()); err != nil {
		log.Fatal(err)
	}
}
