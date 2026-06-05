package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ContentItem struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type InputMessage struct {
	Role    string        `json:"role"`
	Content []ContentItem `json:"content"`
}

type RequestBody struct {
	Model string         `json:"model"`
	Input []InputMessage `json:"input"`
}

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}

	systemPrompt := "You are an experienced software engineering instructor.Provide concise, accurate explanations."

	userPrompt := "Explain what REST API is in 3 sentences."

	body := RequestBody{
		Model: "gpt-5",
		Input: []InputMessage{
			{
				Role: "system",
				Content: []ContentItem{
					{
						Type: "input_text",
						Text: systemPrompt,
					},
				},
			},
			{
				Role: "user",
				Content: []ContentItem{
					{
						Type: "input_text",
						Text: userPrompt,
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/responses",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println(string(responseBody))
}
