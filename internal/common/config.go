package common

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the application configuration
type Config struct {
	// Server configuration
	Port    string `envconfig:"PORT" default:"8080"`
	GinMode string `envconfig:"GIN_MODE" default:"debug"`

	// Database configuration
	DataDir string `envconfig:"DATA_DIR" default:"data"`
	DBDir   string `envconfig:"DB_DIR" default:"db"`

	// Table names
	ProfilesTable string `envconfig:"PROFILES_TABLE" default:"profiles"`
}

// GetPort returns the server port
func (c *Config) GetPort() string {
	return c.Port
}

// GetGinMode returns the Gin mode
func (c *Config) GetGinMode() string {
	return c.GinMode
}

// GetProfilesTablePath returns the full path to the profiles CSV file
func (c *Config) GetProfilesTablePath() string {
	return filepath.Join(c.DataDir, c.DBDir, c.ProfilesTable+".csv")
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	var cfg Config

	// Load from environment variables
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	// Ensure data directories exist
	if err := os.MkdirAll(filepath.Join(cfg.DataDir, cfg.DBDir), 0755); err != nil {
		return nil, err
	}

	return &cfg, nil
}
