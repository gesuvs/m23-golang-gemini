package gemini_setup

import (
	"context"
	"gemini/example/internal/config"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func PromptGemini(prompt string) *genai.GenerateContentResponse {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	apiKey := config.GoDotEnvVariable("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	return resp

}
