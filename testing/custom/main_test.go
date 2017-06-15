package main

import (
	"errors"
	"reflect"
	"testing"
)

type proc struct {
	out []byte
	err error
}

type mockCommand struct {
	ls  proc
	pwd proc
}

func (p mockCommand) Exec(c string) ([]byte, error) {
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
		name   string
		mock   mockCommand
		expOut string
		expErr error
	}{
		{
			"WithoutErrors",
			mockCommand{
				proc{[]byte("abb.go\nbar.go\n"), nil},
				proc{[]byte("boo/bar"), nil}},
			"'a'-'b'", nil,
		}, {
			"WithErrorLs",
			mockCommand{
				proc{nil, errors.New("command not found")},
				proc{[]byte("boo/bar"), nil}},
			"", errors.New("error ls: command not found"),
		}, {
			"WithErrorPwd",
			mockCommand{
				proc{[]byte("abb.go\nbar.go\n"), nil},
				proc{nil, errors.New("command not found")},
			},
			"", errors.New("error pwd: command not found"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(st *testing.T) {
			out, err := process(c.mock)
			if !reflect.DeepEqual(err, c.expErr) {
				st.Fatalf("Expect err to be %q got %q", c.expErr, err)
			}

			if string(c.expOut) != string(out) {
				st.Fatalf("Expected %q got %q", string(c.expOut), string(out))
			}
		})
	}
}
