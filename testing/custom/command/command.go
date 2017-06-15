package command

import (
	"fmt"
	"os/exec"
)

// Commander ...
type Commander interface {
	Exec(string) ([]byte, error)
}

// Command ...
type Command struct{}

// Exec ...
func (c Command) Exec(command string) ([]byte, error) {
	return exec.Command(command).CombinedOutput()
}

// Ls ...
func Ls(c Commander) ([]byte, error) {
	out, err := c.Exec("ls")
	if err != nil {
		return nil, fmt.Errorf("error ls: %s", err)
	}
	return out, nil
}

// Pwd ...
func Pwd(c Commander) ([]byte, error) {
	out, err := c.Exec("pwd")
	if err != nil {
		return nil, fmt.Errorf("error pwd: %s", err)
	}
	return out, nil
}
