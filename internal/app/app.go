package app

import (
	"log"

	"github.com/AlixPa/dummy-backend-flo/internal/common"
)

func Run() {

	cfg := common.LoadConfig()
	router := SetupRouter()

	log.Printf("Server starting on port %s...\n", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
