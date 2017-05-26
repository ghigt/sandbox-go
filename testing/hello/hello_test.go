package main

import "testing"

func TestHello(t *testing.T) {
	exp := "hello, testing"
	res := hello()

	if res != exp {
		t.Fatalf("Expected %q, got %q", exp, res)
	}
}
