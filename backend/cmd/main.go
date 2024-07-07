package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	handlers "github.com/highonsemicolon/cloud-shell/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	logger := logrus.New()

	handler := handlers.NewHandler(logger)

	r.POST("/api/start", handler.Start)
	r.POST("/api/stop", handler.Stop)

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	handler.StopAll()
	logger.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server exiting")
}
