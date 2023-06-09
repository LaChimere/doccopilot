package client

import (
	"context"
	"log"

	"github.com/LaChimere/doccopilot/global"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	ctx    context.Context
	client *openai.Client
}

func NewClient(ctx context.Context) Client {
	client := Client{ctx: ctx}

	if len(global.API_TYPE) == 0 || global.API_TYPE == openai.APITypeOpenAI {
		client.client = openai.NewClient(global.API_KEY)
	}

	if global.API_TYPE == openai.APITypeAzure {
		config := openai.DefaultAzureConfig(global.API_KEY, global.API_ENDPOINT)
		if len(global.API_VERSION) > 0 {
			config.APIVersion = global.API_VERSION
		}

		client.client = openai.NewClientWithConfig(config)
	}

	return client
}

func (c *Client) CreateCompletion(req openai.CompletionRequest) (string, error) {
	resp, err := c.client.CreateCompletion(c.ctx, req)
	if err != nil {
		log.Printf("c.CreateCompletion error: %v", err)
		return "", err
	}

	return resp.Choices[0].Text, nil
}
