package main

import (
	"context"
	"errors"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

const openaiRequestTimeout = time.Second * 10

func goDotEnvVariable(key string) string {
	err := godotenv.Load("./../openai-service/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func (app *Config) CreatePrompt(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Prompt string `json:"prompt"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, errors.New("prompt is invalid"), http.StatusBadRequest)
	}

	ctx := context.Background()

	client := gpt3.NewClient(goDotEnvVariable("OPEN_AI_API_KEY"))

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{requestPayload.Prompt},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Panic("Open AI request failed", err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Open AI request successfull",
		Data:    resp.Choices[0].Text,
	}

	app.writeJson(w, http.StatusAccepted, payload)
}
