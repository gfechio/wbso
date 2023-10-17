package cmd

import "fmt"

func Setup(config string) {
	defaultConfigPath := "~./.wbso/config"
	fmt.Println("Pass binary location (/usr/local/bin)")
	// if read config location and file not there or else read file.
	if config != defaultConfigPath {
		fmt.Println("Pass config.yaml location (%s)", defaultConfigPath)
	}
	// if different location update config
}
