package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	YouTrack YouTrackConfig `yaml:"youtrack"`
	// Add other configurations as needed
}

type YouTrackConfig struct {
	BaseURL   string `yaml:"base_url"`
	Token     string `yaml:"token"`
	ProjectID string `yaml:"project_id"`
}

func ReadConfig(filename string) (AppConfig, error) {
	var config AppConfig

	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
