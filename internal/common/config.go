package common

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
)

type DbTables struct {
	Profiles string
}

type DbTablesCsv struct {
	Profiles string
}
type Config struct {
	// One should not modify the Port value by himself, use config.GetPort() instead
	Port            string      `envconfig:"PORT" default:"8080"`
	rootPath        string      `ignored:"true"`
	dataPath        string      `ignored:"true"`
	dbPath          string      `ignored:"true"`
	dbTablesName    DbTables    `ignored:"true"`
	dbTablesCsvPath DbTablesCsv `ignored:"true"`
}

func (c *Config) GetDbTablesCsvPath() DbTablesCsv {
	return c.dbTablesCsvPath
}

func (c *Config) GetPort() string {
	return c.Port
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	cfg.rootPath = rootPath
	cfg.dataPath = filepath.Join(rootPath, "data")
	cfg.dbPath = filepath.Join(cfg.dataPath, "db")

	cfg.dbTablesName = DbTables{Profiles: "profiles"}
	cfg.dbTablesCsvPath = DbTablesCsv{Profiles: filepath.Join(cfg.dbPath, cfg.dbTablesName.Profiles+".csv")}

	return &cfg, nil
}
