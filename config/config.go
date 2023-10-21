package config

import (
	"github.com/spf13/viper"
)

type TicketManagement struct {
	IssueTracker string `yaml:"issue_tracker"`
	Config IssueTrackerConfig `yaml:"config"`
}

type IssueTrackerConfig struct {
	BaseURL   string `yaml:"base_url"`
	Token     string `yaml:"token"`
	ProjectID string `yaml:"project_id"`
}

func ReadConfig(filename string) (TicketManagement, error) {
	var config TicketManagement

	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
