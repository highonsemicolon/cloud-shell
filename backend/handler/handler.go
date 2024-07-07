package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/highonsemicolon/cloud-shell/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	service services.Service
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{
		logger:  logger,
		service: services.NewService(),
	}
}

func (h *Handler) Start(c *gin.Context) {

	containerID, err := h.service.StartShellInDocker()
	if err != nil {
		h.logger.Error("Failed to start shell", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start shell"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shell started successfully", "containerID": containerID})
}
