package util

import (
	"errors"
	"reflect"
	"testing"
)

type mockCommander struct {
	out []byte
	err error
}

func (m mockCommander) Command(c string) ([]byte, error) {
	return m.out, m.err
}

func TestLs(t *testing.T) {

	cases := []struct {
		mock   mockCommander
		expOut []byte
		expErr error
	}{
		{mockCommander{[]byte("foo.go\nbar.go\n"), nil}, []byte("foo.go\nbar.go\n"), nil},
		{mockCommander{nil, errors.New("command not found: ls")}, nil, errors.New("Error ls: command not found: ls")},
	}

	for _, c := range cases {
		Mock(c.mock)

		out, err := Ls()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Fatalf("Expect err to be %q got %q", c.expErr, err)
		}

		if string(c.expOut) != string(out) {
			t.Fatalf("Expected %q got %q", string(c.expOut), string(out))
		}
	}
}

func TestPwd(t *testing.T) {
	cases := []struct {
		mock   mockCommander
		expOut []byte
		expErr error
	}{
		{mockCommander{[]byte("/foo/bar"), nil}, []byte("/foo/bar"), nil},
		{mockCommander{nil, errors.New("command not found")}, nil, errors.New("Error pwd: command not found")},
	}

	for _, c := range cases {
		Mock(c.mock)

		out, err := Pwd()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Fatalf("Expect err to be %q got %q", c.expErr, err)
		}

		if string(c.expOut) != string(out) {
			t.Fatalf("Expected %q got %q", string(c.expOut), string(out))
		}
	}
}
