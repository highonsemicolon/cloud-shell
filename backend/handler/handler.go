package handlers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger *logrus.Logger
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Start(c *gin.Context) {

	containerID, err := startShellInDocker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start shell"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shell started successfully", "containerID": containerID})
}

func startShellInDocker() (string, error) {
	cmd := exec.Command("sh", "scripts/start.sh")
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	containerID := string(stdout)
	return containerID, nil
}
