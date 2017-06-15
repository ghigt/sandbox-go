package main

import (
	"fmt"
	"os"

	"github.com/ghigt/sandbox-go/testing/custom/command"
)

func process(c command.Commander) (string, error) {
	ls, err := command.Ls(c)
	if err != nil {
		return "", err
	}

	pwd, err := command.Pwd(c)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%q-%q", string(ls)[0], string(pwd)[0]), nil
}

func main() {
	c := command.Command{}
	p, err := process(c)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(p)
}
