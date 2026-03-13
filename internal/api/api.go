package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"bytes"
	"thedekk/AIReview/internal/config"
	"thedekk/AIReview/internal/env"
)


func Request(fileFormatOut, cod, SupplementationPromtString string, CustomPromt bool) (*string, error) {

	client := &http.Client{}

	jsonData, err := os.Open("internal/api/promt.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return nil, err
	}

	defer jsonData.Close()

	var requestBody RequestBody
	if err := json.NewDecoder(jsonData).Decode(&requestBody); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	configJSON, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return nil, err
	}


	if CustomPromt {
		requestBody.Messages[0].Content += configJSON.CustomPromt + " " + SupplementationPromtString + " " + fmt.Sprintf("YOU HAVE TO Answer the %s Language. REALLY IMPORTANT", configJSON.Language)
		requestBody.Messages[0].Content += fmt.Sprintf(" ALSO, THE FILE FORMAT IS %s. REALLY IMPORTANT", fileFormatOut)
	} else {
			requestBody.Messages[0].Content += configJSON.Promt + " " + SupplementationPromtString + " " + fmt.Sprintf("YOU HAVE TO Answer the %s Language. REALLY IMPORTANT", configJSON.Language)
		  requestBody.Messages[0].Content += fmt.Sprintf(" ALSO, THE FILE FORMAT IS %s. REALLY IMPORTANT", fileFormatOut)
	}


	requestBody.Messages[0].Content += cod

	if len(requestBody.Messages[0].Content) > 262144  {
		fmt.Println("Content is too long, truncating to 262144 characters.")
		requestBody.Messages[0].Content = requestBody.Messages[0].Content[:262144]
	}

	resultJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	req, err := http.NewRequest(
		"POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(resultJSON),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	env := env.Config{}
	env.Load()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+env.KeyAPI)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	var Answer Answer

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&Answer); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
 	if len(Answer.Choices) == 0 {
		fmt.Println("No choices in the response")
		return nil, fmt.Errorf("no choices in the response")
	}


	return &Answer.Choices[0].Message.Content, nil
}
