package client

import (
	"errors"

	"github.com/sashabaranov/go-openai"
)

const (
	SummarizerPrefix = "Your task is to create a concise and informative summary that captures the main points and key details of the following document. The summaries should be in your own words and not simply copy-pasting the original text. Please keep the summary length no more than 200 words. The document content is: \n\n"
)

func (c *Client) SummarizePdf(req *openai.CompletionRequest) (string, error) {
	prompt, ok := req.Prompt.(string)
	if !ok {
		return "", errors.New("the prompt type is not string")
	}

	req.Prompt = SummarizerPrefix + prompt

	resp, err := c.CreateCompletion(req)
	if err != nil {
		return "", err
	}

	return resp, nil
}
