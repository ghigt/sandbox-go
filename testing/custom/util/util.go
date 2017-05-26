package util

import (
	"fmt"
	"os/exec"
)

// Commander ...
type Commander interface {
	Command(string) ([]byte, error)
}

// RealCommander ...
type RealCommander struct{}

var commander Commander

// Command ...
func (c RealCommander) Command(command string) ([]byte, error) {
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

func init() {
	commander = RealCommander{}
}
