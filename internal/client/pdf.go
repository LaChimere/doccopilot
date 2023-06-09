package client

import "github.com/sashabaranov/go-openai"

type PdfCompletionRequest struct {
	BasicParameter

	Prompt string `json:"prompt,omitempty" form:"prompt"`
}

const (
	SummarizerPrefix = "Your task is to create a concise and informative summary that captures the main points and key details of the following document. The summaries should be in your own words and not simply copy-pasting the original text. Please keep the summary length no more than 200 words. The document content is: \n\n"
)

func (c *Client) SummarizePdf(req *PdfCompletionRequest) (string, error) {
	completionRequest := openai.CompletionRequest{
		Model:            req.Model,
		Prompt:           SummarizerPrefix + req.Prompt,
		Temperature:      req.Temperature,
		MaxTokens:        req.MaxTokens,
		TopP:             req.TopP,
		FrequencyPenalty: req.FrequencyPenalty,
		PresencePenalty:  req.PresencePenalty,
	}

	resp, err := c.CreateCompletion(completionRequest)
	if err != nil {
		return "", err
	}

	return resp, nil
}
