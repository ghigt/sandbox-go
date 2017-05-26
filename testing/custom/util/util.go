package util

import (
	"fmt"
	"os/exec"
)

// Commander ...
type Commander interface {
	Command(string) ([]byte, error)
}

type realCommander struct{}

var commander Commander = realCommander{}

// Command ...
func (c realCommander) Command(command string) ([]byte, error) {
	return exec.Command(command).CombinedOutput()
}

// Ls ...
func Ls() ([]byte, error) {
	out, err := commander.Command("ls")
	if err != nil {
		return nil, fmt.Errorf("Error ls: %s", err)
	}
	return out, nil
}

// Pwd ...
func Pwd() ([]byte, error) {
	out, err := commander.Command("pwd")
	if err != nil {
		return nil, fmt.Errorf("Error pwd: %s", err)
	}
	return out, nil
}

// Custom ...
func Custom(c Commander) {
	commander = c
}
