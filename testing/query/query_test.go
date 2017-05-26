package main

import (
	"errors"
	"reflect"
	"testing"
)

type MockReleaseInfo struct {
	Tag string
	Err error
}

func (m MockReleaseInfo) GetLatestReleaseTag(repo string) (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.Tag, nil
}

func TestGetReleaseTagMessage(t *testing.T) {
	cases := []struct {
		mock   MockReleaseInfo
		repo   string
		expMsg string
		expErr error
	}{
		{MockReleaseInfo{"v0.1.0", nil}, "dev/null", "The latest release is \"v0.1.0\"", nil},
		{MockReleaseInfo{"v0.1.0", errors.New("TCP timeout")}, "dev/null", "", errors.New("Error querying Github API: TCP timeout")},
	}

	for _, c := range cases {
		msg, err := getReleaseTagMessage(c.mock, c.repo)
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expErr, err)
		}

		if c.expMsg != msg {
			t.Errorf("Expected %q but go %q", c.expMsg, msg)
		}
	}
}

// func TestGetReleaseTagMessage(t *testing.T) {
// 	m := MockReleaseInfo{"v0.1.0", nil}
//
// 	exp := "The latest release is \"v0.1.0\""
// 	msg, err := getReleaseTagMessage(m, "dev/null")
// 	if err != nil {
// 		t.Fatalf("Expected err to be nil but it was %s", err)
// 	}
//
// 	if exp != msg {
// 		t.Fatalf("Expected %q but got %q", exp, msg)
// 	}
// }
