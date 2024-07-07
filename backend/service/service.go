package services

import (
	"os/exec"
	"strings"
)

type Service interface {
	StartShellInDocker() (string, error)
	StopShellInDocker(string) (string, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s service) StartShellInDocker() (string, error) {
	cmd := exec.Command("sh", "scripts/start.sh")
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
