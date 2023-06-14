package app

import (
	"net/http"

	error2 "github.com/LaChimere/doccopilot/pkg/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

// NewResponse creates a new Response instance.
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponse converts the data to a response with the status code http.StatusOK.
func (resp *Response) ToResponse(data any) {
	if data == nil {
		data = gin.H{}
	}

	resp.Ctx.JSON(http.StatusOK, data)
}

// ToErrorResponse converts the error to a response with the status code of the error.
func (resp *Response) ToErrorResponse(err *error2.Error) {
	response := gin.H{
		"code":    err.Code(),
		"message": err.Message(),
	}

	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	resp.Ctx.JSON(err.StatusCode(), response)
}
