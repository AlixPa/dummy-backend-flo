package common

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	RootPath string
	DataPath string
	DbPath   string
}

func LoadConfig() *Config {
	loadEnv()
	root, _ := os.Getwd()

	config := &Config{
		Port:     getEnv("PORT", "8080"),
		RootPath: root,
		DataPath: path.Join(root, "data"),
		DbPath:   path.Join(root, "data", "db"),
	}

	return config
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
