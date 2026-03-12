package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigJSON struct {
	Promt       string `json:"Promt"`
	Language    string `json:"Language"`
	CustomPromt string `json:"CustomPromt"`
}

func LoadConfig() (ConfigJSON, error) {
	var config ConfigJSON

	// Open
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return config, err
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Error decoding config file:", err)
		return config, err
	}

	return config, nil
}

