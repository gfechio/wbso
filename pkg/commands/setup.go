package cmd

import (
	"fmt"
)

// Constants for default paths
const (
	DefaultConfigPath = "~/.wbso/config"
	DefaultBinPath    = "/usr/local/bin"
)

// Setup initializes the configuration.
func Setup(configPath string) {
	fmt.Printf("Pass binary location (%s): ", DefaultBinPath)
	binLocation := readInput(DefaultBinPath)

	if configPath != DefaultConfigPath {
		fmt.Printf("Pass config.yaml location (%s): ", DefaultConfigPath)
		configLocation := readInput(DefaultConfigPath)

		// Handle configuration updates if different location provided
		updateConfig(configLocation)
	}

	// Now you can use 'binLocation' and 'configPath' as needed.
	fmt.Printf("Binary Location: %s\n", binLocation)
	fmt.Printf("Config Location: %s\n", configPath)
}

// readInput reads user input or returns the default value.
func readInput(defaultValue string) string {
	var input string
	fmt.Scanln(&input)
	if input == "" {
		return defaultValue
	}
	return input
}

// updateConfig simulates updating the configuration.
func updateConfig(configLocation string) {
	fmt.Printf("Updating config using location: %s\n", configLocation)
	// Implement your configuration update logic here
}
