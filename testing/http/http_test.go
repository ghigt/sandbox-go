package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Root request error: %s", err)
	}

	cases := []struct {
		w       *httptest.ResponseRecorder
		r       *http.Request
		token   string
		expCode int
		expBody []byte
		expLogs []string
	}{
		{httptest.NewRecorder(), req, "magic", http.StatusOK, []byte("You have some magic in you"), []string{"Allowed an access attempt\n"}},
		{httptest.NewRecorder(), req, "", http.StatusForbidden, []byte("You don't have enough magic in you\n"), []string{"Denied an access attempt\n"}},
	}

	for _, c := range cases {
		rd, wr := io.Pipe()
		buf := bufio.NewReader(rd)

		log.SetOutput(wr)

		c.r.Header.Set("X-Access-Token", c.token)

		go func() {
			for _, expLine := range c.expLogs {
				msg, err := buf.ReadString('\n')
				if err != nil {
					t.Errorf("expected to be able to read from log but got error: %s", err)
				}

				if !strings.HasSuffix(msg, expLine) {
					t.Errorf("Log line didn't match suffix:\n\t%q\n\t%q", expLine, msg)
				}
			}
		}()

		mainHandler(c.w, c.r)

		if c.expCode != c.w.Code {
			t.Errorf("Status Code didn't match:\n\t%q\n\t%q", c.expCode, c.w.Code)
		}

		if !bytes.Equal(c.expBody, c.w.Body.Bytes()) {
			t.Errorf("Body didn't match:\n\t%q\n\t%q", string(c.expBody), c.w.Body.String())
		}
	}
}
