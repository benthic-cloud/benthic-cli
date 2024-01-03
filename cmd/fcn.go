package cmd

import (
	"benthic/config"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var fcnCmd = &cobra.Command{
	Use:   "fcn",
	Short: "Commands to interact with fcn",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a function",
	Run:   run,
}

// Define a shared HTTP client
var client = &http.Client{}

// Function to make a POST request using API client
func makePostRequest(url, apiKey string, body []byte) (*http.Response, error) {
	requestBody := bytes.NewBuffer(body)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return nil, err
	}

	// Set the necessary headers
	req.Header.Set("apikey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform the request using the shared client
	return client.Do(req)
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Running function...")

	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return // Exit the function if config loading fails
	}

	// Extracting the endpoint and API key from the loaded config
	endpoint := cfg.Endpoint         // Make sure your Config struct has an "Endpoint" field
	apiKey := cfg.Profiles[0].APIKey // Assuming you want to use the first profile

	// Define the specific API path you want to hit
	apiPath := "fcn/start"
	fullURL := endpoint + apiPath

	// Assuming an empty JSON body '{}' for the POST request
	body := []byte("{}")

	// Make the POST request
	response, err := makePostRequest(fullURL, apiKey, body)
	if err != nil {
		fmt.Printf("Error making the request: %v\n", err)
		return // Exit the function if the request fails
	}
	defer response.Body.Close()

	// Handle the response...
	fmt.Printf("Response Status: %s\n", response.Status)
}

func init() {
	fcnCmd.AddCommand(runCmd)
	RootCmd.AddCommand(fcnCmd)
}
