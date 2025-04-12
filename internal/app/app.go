package app

import (
	"log"

	"github.com/AlixPa/dummy-backend-flo/internal/common"
)

func Run() {

	cfg, err := common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	router := SetupRouter(cfg)
	log.Print(cfg.GetPort())
	if err := router.Run(":" + cfg.GetPort()); err != nil {
		log.Fatal(err)
	}
}
