package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestStatusIssuesCmd(t *testing.T) {
	buff := new(bytes.Buffer)
	oldClient := client
	defer func() { client = oldClient }()
	client = newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "https://api.github.com/repos/TheAlgorithms/Go/issues/1" {
			t.Errorf("ListIssues URL = %v, want %v", req.URL, "http://api.github.com/repos/TheAlgorithms/Go/issues/1")
		}
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
	"number": 1,
	"title": "Test Issue 1",
	"body": "Test Issue 1 body",
	"html_url": "https://sample.url",
	"state": "open"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	rootCmd.SetOut(buff)
	rootCmd.SetArgs([]string{"status", "issue", "-r", "TheAlgorithms/Go", "-n", "1"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	got := buff.String()
	want := "\x1b[32mStatus: open\n\x1b[32mTitle: Test Issue 1\n\x1b[32mURL: https://sample.url\n\x1b[32mNumber: 1\n\x1b[32mBody: Test Issue 1 body\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
