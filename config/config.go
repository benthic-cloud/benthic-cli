package config

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
