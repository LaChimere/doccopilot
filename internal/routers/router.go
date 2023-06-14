package routers

import (
	v1 "github.com/LaChimere/doccopilot/internal/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Service is alive
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiV1 := router.Group("/api/v1")
	apiV1.Use()

	// PDF APIs
	pdf := v1.NewPdf()
	{
		apiV1.POST("/pdf/summarize", pdf.Summarize)
	}

	return router
}
