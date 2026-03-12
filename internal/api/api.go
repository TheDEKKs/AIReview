package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"thedekk/AIReview/internal/config"
	"thedekk/AIReview/internal/env"
)

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type Reasoning struct {
    Enabled bool `json:"enabled"`
}
type RequestBody struct {
    Model     string    `json:"model"`
    Messages  []Message `json:"messages"`
    Reasoning Reasoning `json:"reasoning"`
}

func Request(cod, SupplementationPromtString string, CustomPromt bool) error {
	client := &http.Client{}

	jsonData, err := os.Open("internal/api/test.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return err
	}

	defer jsonData.Close()

	var requestBody RequestBody
	if err := json.NewDecoder(jsonData).Decode(&requestBody); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	configJSON, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return err
	}

	requestBody.Messages[0].Content += cod

	if CustomPromt {
		requestBody.Messages[0].Content += configJSON.CustomPromt + " " + SupplementationPromtString + " " + fmt.Sprintf("Reply in - Language %s", configJSON.Language)
	} else {
			requestBody.Messages[0].Content += configJSON.Promt + " " +  fmt.Sprintf("Reply in - Language %s", configJSON.Language) + " " + SupplementationPromtString
	}
	req, err := http.NewRequest(
		"POST", "https://openrouter.ai/api/v1/chat/completions", jsonData,
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	env := env.Config{}
	env.Load()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+env.KeyAPI)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	fmt.Println(string(body))
	return nil
}
