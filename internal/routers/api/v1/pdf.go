package v1

import (
	"log"
	"strings"

	client2 "github.com/LaChimere/doccopilot/internal/client"
	"github.com/LaChimere/doccopilot/pkg/app"
	error2 "github.com/LaChimere/doccopilot/pkg/error"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type Pdf struct{}

func NewPdf() Pdf {
	return Pdf{}
}

//	@title			Summarize PDF contents
//	@version		v1
//	@description	Summarize PDF contents
//	@BasePath		/api/v1
//	@Router			/pdf/summarize [post]
//	@Produce		json
//	@Param			request	body	map[string]any	true				"The OpenAI chat completion request, see https://platform.openai.com/docs/api-reference/chat/create"
//	@Success		200		string	json			"{"summary": "The	summary	of	the	content"}"
func (p Pdf) Summarize(c *gin.Context) {
	client := client2.NewClient(c.Request.Context())
	response := app.NewResponse(c)

	param := openai.ChatCompletionRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		e := error2.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(e)
		return
	}

	if len(param.Messages) == 0 || len(param.Messages[0].Content) == 0 {
		e := error2.InvalidParams.WithDetails("the content is empty")
		response.ToErrorResponse(e)
		return
	}

	data, err := client.SummarizePdf(&param)
	if err != nil {
		log.Printf("client.SummarizePdf error: %v", err)
		e := error2.SummarizePdfFail.WithDetails(err.Error())
		response.ToErrorResponse(e)
		return
	}

	response.ToResponse(map[string]string{
		"summary": strings.TrimSpace(data),
	})
}
