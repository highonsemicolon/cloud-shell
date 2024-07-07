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

func (h *Handler) Stop(c *gin.Context) {

	var body struct {
		ContainerID string `json:"containerID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	containerID, err := h.service.StopShellInDocker(body.ContainerID)
	if err != nil {
		h.logger.Error("Failed to stop shell", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop shell"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shell stopped successfully", "containerID": containerID})
}
