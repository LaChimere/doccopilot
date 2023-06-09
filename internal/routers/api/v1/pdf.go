package v1

import (
	"log"
	"strings"

	client2 "github.com/LaChimere/doccopilot/internal/client"
	"github.com/LaChimere/doccopilot/pkg/app"
	error2 "github.com/LaChimere/doccopilot/pkg/error"

	"github.com/gin-gonic/gin"
)

type Pdf struct{}

func NewPdf() Pdf {
	return Pdf{}
}

func (p Pdf) Summarize(c *gin.Context) {
	client := client2.NewClient(c.Request.Context())
	response := app.NewResponse(c)

	param := client2.PdfCompletionRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		e := error2.InvalidParams.WithDetails(err.Error())
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
