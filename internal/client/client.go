package client

import (
	"context"
	"log"
	"strings"

	"github.com/LaChimere/doccopilot/global"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	ctx    context.Context
	client *openai.Client
}

// NewClient creates a new OpenAI client according to the API type.
func NewClient(ctx context.Context) *Client {
	client := &Client{ctx: ctx}

	if len(global.API_TYPE) == 0 || strings.EqualFold(global.API_TYPE, string(openai.APITypeOpenAI)) {
		client.client = openai.NewClient(global.API_KEY)
	}

	if strings.EqualFold(global.API_TYPE, string(openai.APITypeAzure)) {
		config := openai.DefaultAzureConfig(global.API_KEY, global.API_ENDPOINT)
		if len(global.API_VERSION) > 0 {
			config.APIVersion = global.API_VERSION
		}

		client.client = openai.NewClientWithConfig(config)
	}

	return client
}

// CreateChatCompletion creates a chat completion and send the request to the OpenAI service.
func (c *Client) CreateChatCompletion(req *openai.ChatCompletionRequest) (string, error) {
	resp, err := c.client.CreateChatCompletion(c.ctx, *req)
	if err != nil {
		log.Printf("c.CreateChatCompletion error: %v", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
