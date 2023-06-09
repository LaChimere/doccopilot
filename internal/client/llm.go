package client

import "github.com/sashabaranov/go-openai"

type BasicParameter struct {
	Model            string  `json:"model" form:"model" binding:"required"`
	Temperature      float32 `json:"temperature,omitempty" form:"temperature" binding:"min=0.0,max=1.0"`
	MaxTokens        int     `json:"max_tokens,omitempty" form:"max_tokens" binding:"min=0,max=2048"`
	TopP             float32 `json:"top_p,omitempty" form:"top_p" binding:"min=0.0,max=1.0"`
	FrequencyPenalty float32 `json:"frequency_penalty,omitempty" form:"frequency_penalty" binding:"min=0.0,max=2.0"`
	PresencePenalty  float32 `json:"presence_penalty,omitempty" form:"presence_penalty" binding:"min=0.0,max=2.0"`
}

func DefaultCompletionRequest() BasicParameter {
	return BasicParameter{
		Model:            openai.GPT3TextDavinci003,
		Temperature:      0.7,
		MaxTokens:        256,
		TopP:             1.0,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
	}
}
