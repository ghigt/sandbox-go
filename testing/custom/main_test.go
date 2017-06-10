package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ghigt/sandbox-go/testing/custom/util"
)

type command struct {
	out []byte
	err error
}

type processCommand struct {
	ls  command
	pwd command
}

func (p processCommand) Command(c string) ([]byte, error) {
	switch c {
	case "ls":
		return p.ls.out, p.ls.err
	case "pwd":
		return p.pwd.out, p.pwd.err
	}
	return nil, nil
}

func TestProcess(t *testing.T) {
	cases := []struct {
		command processCommand
		expOut  string
		expErr  error
	}{
		{
			processCommand{
				command{[]byte("abb.go\nbar.go\n"), nil},
				command{[]byte("boo/bar"), nil}},
			"'a'-'b'", nil,
		}, {
			processCommand{
				command{nil, errors.New("command not found")},
				command{[]byte("boo/bar"), nil}},
			"", errors.New("Error ls: command not found"),
		}, {
			processCommand{
				command{[]byte("abb.go\nbar.go\n"), nil},
				command{nil, errors.New("command not found")},
			},
			"", errors.New("Error pwd: command not found"),
		},
	}

	for _, c := range cases {
		util.Custom(c.command)

		out, err := process()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Fatalf("Expect err to be %q got %q", c.expErr, err)
		}

		if string(c.expOut) != string(out) {
			t.Fatalf("Expected %q got %q", string(c.expOut), string(out))
		}
	}
}
