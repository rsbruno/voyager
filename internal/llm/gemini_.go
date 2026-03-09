package llm

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
	model  string
}

func NewGeminiClient() (*GeminiClient, error) {

	ctx := context.Background()

	c, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("llm: erro ao criar client: %w", err)
	}

	return &GeminiClient{
		client: c,
		model:  "gemini-3-flash-preview",
	}, nil
}

func (c *GeminiClient) Execute(prompt string) (string, error) {

	if prompt == "" {
		return "", fmt.Errorf("llm: prompt vazio")
	}

	ctx := context.Background()

	result, err := c.client.Models.GenerateContent(
		ctx,
		c.model,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("llm: erro ao executar modelo: %w", err)
	}

	return result.Text(), nil
}