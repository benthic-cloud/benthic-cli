package cmd

import (
	"benthic/config"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update and manage configuration",
}

var defaultNewConfig = config.Config{
	Profiles: []config.Profile{
		{
			Name:   "default",
			APIKey: "REPLACE_ME",
		},
	},
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return !os.IsNotExist(err) && !info.IsDir()
}

func writeToYAML(c config.Config, filename string) {
	if fileExists(filename) {
		fmt.Println("File already exists.")
		return
	}

	yamlData, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = os.WriteFile(filename, yamlData, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("Data written to file successfully.")
}

func initConfig(cmd *cobra.Command, args []string) {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find the home directory:", err)
	}

	// Construct the new path
	newPath := filepath.Join(homeDir, ".benthic")

	// Create the directory with os.ModePerm permissions
	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create the directory:", err)
	} else {
		log.Println("Directory created successfully:", newPath)
	}

	// If the config file doesn't exist at ~/.benthic/config.yml, create it using the defaultNewConfig struct.
	newPath = filepath.Join(newPath, "config.yml")

	writeToYAML(defaultNewConfig, newPath)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration file at ~/.benthic/config.yml",
	Run:   initConfig,
}

func init() {
	configCmd.AddCommand(initCmd)
	RootCmd.AddCommand(configCmd)
}