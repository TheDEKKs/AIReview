package api

import (
	"net/http"
	"io"
	"os"
	"fmt"
	"thedekk/AIReview/internal/env"
)

func Test() error {
	client := &http.Client{}

	jsonData, err := os.Open("internal/api/test.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return err
	}

	defer jsonData.Close()

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
