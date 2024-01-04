package cmd

import (
	"benthic/config"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
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

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get the logs for a function",
	Run:   logs,
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

	image := cmd.Flag("image").Value.String()
	env := cmd.Flag("env").Value.String()

	// Create the JSON body
	body := []byte(fmt.Sprintf(`{"image":"%s", "env":%s}`, image, env))

	// Make the POST request
	response, err := makePostRequest(fullURL, apiKey, body)
	if err != nil {
		fmt.Printf("Error making the request: %v\n", err)
		return // Exit the function if the request fails
	}
	defer response.Body.Close()

	// Read the response body into a string
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	bodyString := buf.String()

	// Handle the response...
	fmt.Printf("%s\n", bodyString)
}

// logs function to handle 'logs' command
func logs(cmd *cobra.Command, args []string) {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return // Exit the function if config loading fails
	}

	functionID := cmd.Flag("function-id").Value.String()

	apiKey := cfg.Profiles[0].APIKey // Assuming you want to use the first profile

	// Define the URL
	url := fmt.Sprintf("%s/fcn/logs/%s", cfg.Endpoint, functionID)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating the request: %s\n", err.Error())
		return
	}

	// Add apikey to the request headers
	req.Header.Set("apikey", apiKey)

	// Send the request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err.Error())
		return
	}
	defer response.Body.Close()

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err.Error())
		return
	}

	// Convert the body to a string and print
	bodyString := string(bodyBytes)
	fmt.Printf("%s\n", bodyString)
}

func init() {
	runCmd.Flags().StringP("image", "i", "", "public Docker image to run")
	runCmd.Flags().StringP("env", "e", "[]", "JSON variables to pass to the function")
	fcnCmd.AddCommand(runCmd)

	logsCmd.Flags().StringP("function-id", "i", "", "ID of the function to get logs for")
	fcnCmd.AddCommand(logsCmd)

	RootCmd.AddCommand(fcnCmd)
}
