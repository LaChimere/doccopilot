package client

import (
	"github.com/sashabaranov/go-openai"
)

const (
	SummarizerPrefix = "Your task is to create a concise and informative summary that captures the main points and key details of the following document. The summaries should be in your own words and not simply copy-pasting the original text. Please keep the summary length no more than 50 words. The document content is below: \n\n"
)

// SummarizePdf is to summarize the pdf content into a short paragraph.
func (c *Client) SummarizePdf(req *openai.ChatCompletionRequest) (string, error) {
	message := SummarizerPrefix + req.Messages[0].Content
	req.Messages[0].Content = message

	resp, err := c.CreateChatCompletion(req)
	if err != nil {
		return "", err
	}

	return resp, nil
}
