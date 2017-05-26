package util

import "testing"

type mockCommander struct{}

func (m mockCommander) CombinedOutput(c string) ([]byte, error) {
	if c == "ls" {
		return []byte("foo.go\nbar.go\n"), nil
	}
	if c == "pwd" {
		return []byte("/foo/bar"), nil
	}
	return []byte{}, nil
}

func TestLs(t *testing.T) {
	Mock(mockCommander{})
	ls, err := Ls()
	if err != nil {
		t.Fatalf("Expect no error got %s", err)
	}

	exp := "foo.go\nbar.go\n"
	if string(ls) != exp {
		t.Fatalf("Expected %q got %q", exp, ls)
	}
}

func TestPwd(t *testing.T) {
	pwd, err := Pwd()
	if err != nil {
		t.Fatalf("Expect no error got %s", err)
	}

	exp := "/foo/bar"
	if string(pwd) != exp {
		t.Fatalf("Expected %q got %q", exp, pwd)
	}
}
