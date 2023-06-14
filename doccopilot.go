package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LaChimere/doccopilot/global"
	"github.com/LaChimere/doccopilot/internal/routers"
)

func init() {
	global.API_KEY = os.Getenv("API_KEY")
	global.API_TYPE = os.Getenv("API_TYPE")
	global.API_VERSION = os.Getenv("API_VERSION")
	global.API_ENDPOINT = os.Getenv("BASE_URL")
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe error: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
