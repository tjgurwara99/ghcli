package api_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/tjgurwara99/ghcli/api"
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

func Test_api_ListIssues(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseUrl string
	}
	client := newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "http://api.github.com/repos/TheAlgorithms/issues?state=open" {
			t.Errorf("ListIssues URL = %v, want %v", req.URL, "http://api.github.com/repos/TheAlgorithms")
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
	type args struct {
		repo  string
		state string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []api.Issue
		wantErr bool
	}{
		{
			name: "Test_api_ListIssues",
			fields: fields{
				client:  client,
				baseUrl: "http://api.github.com/",
			},
			args: args{
				repo:  "TheAlgorithms",
				state: "open",
			},
			want: []api.Issue{
				{
					Number:      1,
					Title:       "Test Issue 1",
					Body:        "Test Issue 1 body",
					URL:         "https://sample.url",
					State:       "open",
					PullRequest: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := api.NewApi(tt.fields.client, tt.fields.baseUrl)
			got, err := a.ListIssues(tt.args.repo, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListIssues() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_api_ListPRs(t *testing.T) {
	t.Parallel()
	type fields struct {
		client  *http.Client
		baseUrl string
	}
	client := newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "http://api.github.com/repos/TheAlgorithms/pulls?state=open" {
			t.Errorf("ListIssues URL = %v, want %v", req.URL, "http://api.github.com/repos/TheAlgorithms")
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
	type args struct {
		repo  string
		state string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []api.PullRequest
		wantErr bool
	}{
		{
			name: "Test_api_ListIssues",
			fields: fields{
				client:  client,
				baseUrl: "http://api.github.com/",
			},
			args: args{
				repo:  "TheAlgorithms",
				state: "open",
			},
			want: []api.PullRequest{
				{
					Number: 1,
					Title:  "Test Issue 1",
					Body:   "Test Issue 1 body",
					URL:    "https://sample.url",
					State:  "open",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := api.NewApi(tt.fields.client, tt.fields.baseUrl)
			got, err := a.ListPRs(tt.args.repo, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListIssues() got = %v, want %v", got, tt.want)
			}
		})
	}
}
