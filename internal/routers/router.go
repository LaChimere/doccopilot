package routers

import (
	"net/http"

	"github.com/LaChimere/doccopilot/docs"
	v1 "github.com/LaChimere/doccopilot/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Service is alive
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Swagger API docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("/api/v1")
	apiV1.Use()

	// PDF APIs
	pdf := v1.NewPdf()
	{
		apiV1.POST("/pdf/summarize", pdf.Summarize)
	}

	return router
}
