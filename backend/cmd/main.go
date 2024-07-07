package main

import (
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

	if err := r.Run("localhost:8080"); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
