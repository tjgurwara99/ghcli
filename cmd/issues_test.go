package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type roundTripFunc func(req *http.Request) *http.Response

func (r roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestListIssuesCmd(t *testing.T) {
	buff := new(bytes.Buffer)
	oldClient := client
	defer func() { client = oldClient }()
	client = newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "https://api.github.com/repos/TheAlgorithms/Go/issues?state=open" {
			t.Errorf("ListIssues URL = %v, want %v", req.URL, "http://api.github.com/repos/TheAlgorithms/Go/issues?state=open")
		}
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`[
	{
		"number": 1,
		"title": "Test Issue 1",
		"body": "Test Issue 1 body",
		"html_url": "https://sample.url",
		"state": "open"
	}
]`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	rootCmd.SetOut(buff)
	rootCmd.SetArgs([]string{"list", "issues", "TheAlgorithms/Go"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	got := buff.String()
	want := "{State:open Body:Test Issue 1 body Title:Test Issue 1 URL:https://sample.url Number:1 PullRequest:<nil>}\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
