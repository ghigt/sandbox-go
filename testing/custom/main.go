package main

import (
	"_sandbox/testing/custom/util"
	"fmt"
	"log"
)

func process() (string, error) {
	ls, err := util.Ls()
	if err != nil {
		return "", err
	}

	pwd, err := util.Pwd()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%q-%q", string(ls)[0], string(pwd)[0]), nil
}

func main() {
	p, err := process()
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(p)
}
