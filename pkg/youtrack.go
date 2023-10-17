package youtrack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/spf13/viper"
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

