package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// Profile represents a profile in the config file
type Profile struct {
	// Name of the profile
	Name string
	// API Key for the profile
	APIKey string
}

// Config represents the config file
type Config struct {
	// Benthic Cloud API endpoint
	Endpoint string

	// List of profiles
	Profiles []Profile
}

func LoadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not find the home directory: %v", err)
	}
	configPath := filepath.Join(home, ".benthic", "config.yaml")
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config file: %v", err)
	}

	return &config, nil
}
