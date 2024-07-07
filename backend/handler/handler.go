package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	services "github.com/highonsemicolon/cloud-shell/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger       *logrus.Logger
	service      services.Service
	containerMap map[string]string

	mu sync.Mutex
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{
		logger:       logger,
		service:      services.NewService(),
		containerMap: make(map[string]string),
		mu:           sync.Mutex{},
	}
}

func (h *Handler) Start(c *gin.Context) {
	sessionID := "shell-" + h.service.GenerateSessionID()
	containerID, err := h.service.StartShellInDocker(sessionID)
	if err != nil {
		h.logger.Error("Failed to start shell", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start shell"})
		return
	}

	h.mu.Lock()
	h.containerMap[sessionID] = containerID
	h.mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Shell started successfully", "sessionID": sessionID})
}

func (h *Handler) Stop(c *gin.Context) {

	var body struct {
		SessionID string `json:"sessionID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	containerID, exists := h.containerMap[body.SessionID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	_, err := h.service.StopShellInDocker(containerID)
	if err != nil {
		h.logger.Error("Failed to stop shell", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stop shell"})
		return
	}

	h.mu.Lock()
	delete(h.containerMap, body.SessionID)
	h.mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Shell stopped successfully", "sessionID": body.SessionID})
}

func (h *Handler) StopAll() {
	for sessionID, containerID := range h.containerMap {
		h.service.StopShellInDocker(containerID)
		delete(h.containerMap, sessionID)
	}
	h.logger.Info("All shells stopped")
}
