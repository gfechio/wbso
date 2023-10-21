package ticketmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type Config struct {
	BaseURL   string `mapstructure:"base_url"`
	Token     string `mapstructure:"token"`
	ProjectID string `mapstructure:"project_id"`
}

type Issue struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	// Add other fields as needed
}

func ReadConfig(filename string) (Config, error) {
	var config Config

	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func NewYouTrackClient(config Config) *http.Client {
	client := &http.Client{}

	return client
}

func GetIssues(client *http.Client, config Config) ([]Issue, error) {
	url := fmt.Sprintf("%s/youtrack/api/issues?project=%s", config.BaseURL, config.ProjectID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+config.Token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err := json.Unmarshal(body, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func UpdateIssueDescription(client *http.Client, config Config, issueID, newDescription string) error {
	// Construct the URL for updating the issue's description
	url := fmt.Sprintf("%s/youtrack/api/issues/%s", config.BaseURL, issueID)

	// Create a map to represent the new description
	updateData := map[string]string{
		"description": newDescription,
	}

	// Marshal the update data to JSON
	updateJSON, err := json.Marshal(updateData)
	if err != nil {
		return err
	}

	// Create a new HTTP PUT request with the updated description JSON
	req, err := http.NewRequest("PUT", url, bytes.NewReader(updateJSON))
	if err != nil {
		return err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Set the Authorization header with the provided token
	req.Header.Set("Authorization", "Bearer "+config.Token)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the update was successful (HTTP status 200 OK or 201 Created)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to update issue description, status code: %d", resp.StatusCode)
	}

	return nil
}
