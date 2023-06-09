package global

import "github.com/sashabaranov/go-openai"

// Global client configs
var (
	API_KEY      string
	API_ENDPOINT string
	API_VERSION  string
	API_TYPE     openai.APIType
)
