package services

import (
	"os/exec"
	"strings"

	"golang.org/x/exp/rand"
)

type Service interface {
	StartShellInDocker(string) (string, error)
	StopShellInDocker(string) (string, error)
	GenerateSessionID() string
}

type service struct {
}

func NewService() Service {
	return &service{}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (s service) StartShellInDocker(sessionID string) (string, error) {
	cmd := exec.Command("sh", "scripts/start.sh", sessionID)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	containerID := strings.TrimSpace(string(stdout))
	return containerID, nil
}

func (s service) StopShellInDocker(id string) (string, error) {
	cmd := exec.Command("sh", "scripts/stop.sh", id)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	containerID := strings.TrimSpace(string(stdout))
	return containerID, nil
}

func (s service) GenerateSessionID() string {
	b := make([]byte, 12)
	for i := range b {
		b[i] = byte(letters[rand.Intn(len(letters))])
	}
	return string(b)
}
