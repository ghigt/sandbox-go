package util

import "os/exec"

// Commander ...
type Commander interface {
	CombinedOutput(string) ([]byte, error)
}

// RealCommander ...
type RealCommander struct{}

var commander Commander

// CombinedOutput ...
func (c RealCommander) CombinedOutput(command string) ([]byte, error) {
	return exec.Command(command).CombinedOutput()
}

// Ls ...
func Ls() ([]byte, error) {
	return commander.CombinedOutput("ls")
}

// Pwd ...
func Pwd() ([]byte, error) {
	return commander.CombinedOutput("pwd")
}

func init() {
	commander = RealCommander{}
}
