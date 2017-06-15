package command

import (
	"errors"
	"reflect"
	"testing"
)

type mockCommander struct {
	out []byte
	err error
}

func (m mockCommander) Exec(c string) ([]byte, error) {
	return m.out, m.err
}

func TestLs(t *testing.T) {

	cases := []struct {
		name   string
		mock   mockCommander
		expOut []byte
		expErr error
	}{
		{"WithoutError", mockCommander{[]byte("foo.go\nbar.go\n"), nil}, []byte("foo.go\nbar.go\n"), nil},
		{"WithError", mockCommander{nil, errors.New("command not found: ls")}, nil, errors.New("error ls: command not found: ls")},
	}

	for _, c := range cases {
		t.Run(c.name, func(st *testing.T) {
			out, err := Ls(c.mock)
			if !reflect.DeepEqual(err, c.expErr) {
				st.Fatalf("Expect err to be %q got %q", c.expErr, err)
			}

			if string(c.expOut) != string(out) {
				st.Fatalf("Expected %q got %q", string(c.expOut), string(out))
			}
		})
	}
}

func TestPwd(t *testing.T) {
	cases := []struct {
		name   string
		mock   mockCommander
		expOut []byte
		expErr error
	}{
		{"WithoutErrors", mockCommander{[]byte("/foo/bar"), nil}, []byte("/foo/bar"), nil},
		{"WithErrors", mockCommander{nil, errors.New("command not found")}, nil, errors.New("error pwd: command not found")},
	}

	for _, c := range cases {
		t.Run(c.name, func(st *testing.T) {
			out, err := Pwd(c.mock)
			if !reflect.DeepEqual(err, c.expErr) {
				st.Fatalf("Expect err to be %q got %q", c.expErr, err)
			}

			if string(c.expOut) != string(out) {
				st.Fatalf("Expected %q got %q", string(c.expOut), string(out))
			}
		})
	}
}
