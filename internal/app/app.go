package app

import (
	"log"

	"github.com/AlixPa/dummy-backend-flo/internal/common"
)

func Run() {

	cfg := common.LoadConfig()
	router := SetupRouter(cfg)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
