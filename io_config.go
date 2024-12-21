package main

import (
	"encoding/json"
	"log"
	"os"

	"example.com/m/models"
)

func loadConfig() models.Config {
	var config models.Config
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Println("Config not found, using defaults")
		return config
	}
	json.Unmarshal(data, &config)
	return config
}

func saveConfig(config models.Config) {
	data, err := json.Marshal(config)
	if err != nil {
		log.Println("Failed to marshal config:", err)
		return
	}
	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		log.Println("Failed to write config file:", err)
	}
}
