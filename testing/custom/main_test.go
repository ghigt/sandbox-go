package main

import (
	"_sandbox/testing/custom/util"
	"testing"
)

type mockCommander struct{}

func (m mockCommander) Command(c string) ([]byte, error) {
	if c == "ls" {
		return []byte("foo.go\nbar.go\n"), nil
	}
	if c == "pwd" {
		return []byte("/foo/bar"), nil
	}
	return []byte{}, nil
}

func TestProcess(t *testing.T) {
	util.Mock(mockCommander{})
	p, err := process()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	exp := "'f'-'/'"
	if p != exp {
		t.Fatalf("Expected %q got %q", exp, p)
	}
}
