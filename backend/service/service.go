package services

import (
	"os/exec"
)

type Service interface {
	StartShellInDocker() (string, error)
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
	containerID := string(stdout)
	return containerID, nil
}
