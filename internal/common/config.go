package common

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	rootPath     string
	dataPath     string
	dbPath       string
	dbTablesName dbTables
	Routeur      *RouteurConfig
}
type dbTables struct {
	profiles string
}

type RouteurConfig struct {
	Profile *ProfileConfig
}

type ProfileConfig struct {
	DbCsvPath string
}

func (cfg *Config) getProfileConfig() *ProfileConfig {
	return &ProfileConfig{
		DbCsvPath: path.Join(cfg.dbPath, cfg.dbTablesName.profiles+".csv"),
	}
}

func (cfg *Config) getRouteurConfig() *RouteurConfig {
	return &RouteurConfig{
		Profile: cfg.getProfileConfig(),
	}
}

func LoadConfig() *Config {
	loadEnv()
	root, _ := os.Getwd()

	config := &Config{
		Port:     getEnv("PORT", "8080"),
		rootPath: root,
		dataPath: path.Join(root, "data"),
		dbPath:   path.Join(root, "data", "db"),
		dbTablesName: dbTables{
			profiles: "profiles",
		},
	}
	config.Routeur = config.getRouteurConfig()

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
