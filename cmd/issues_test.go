package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type roundTripFunc func(req *http.Request) *http.Response

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
//
// RoundTrip should not attempt to interpret the response. In
// particular, RoundTrip must return err == nil if it obtained
// a response, regardless of the response's HTTP status code.
// A non-nil err should be reserved for failure to obtain a
// response. Similarly, RoundTrip should not attempt to
// handle higher-level protocol details such as redirects,
// authentication, or cookies.
//
// RoundTrip should not modify the request, except for
// consuming and closing the Request's Body. RoundTrip may
// read fields of the request in a separate goroutine. Callers
// should not mutate or reuse the request until the Response's
// Body has been closed.
//
// RoundTrip must always close the body, including on errors,
// but depending on the implementation may do so in a separate
// goroutine even after RoundTrip returns. This means that
// callers wanting to reuse the body for subsequent requests
// must arrange to wait for the Close call before doing so.
//
// The Request's URL and Header fields must be initialized.
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
