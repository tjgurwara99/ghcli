package api_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-github/github"
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
		if req.URL.String() != "https://api.github.com/repos/tjgurwara99/Go/issues?state=open" {
			t.Errorf("ListIssues URL = %v, want %v", req.URL, "https://api.github.com/repos/TheAlgorithms")
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
	test1Num := 1
	test1Title := "Test Issue 1"
	test1Body := "Test Issue 1 body"
	test1URL := "https://sample.url"
	test1State := "open"
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*github.Issue
		wantErr bool
	}{
		{
			name: "Test_api_ListIssues",
			fields: fields{
				client:  client,
				baseUrl: "http://api.github.com/",
			},
			args: args{
				repo:  "tjgurwara99/Go",
				state: "open",
			},
			want: []*github.Issue{
				{
					Number:  &test1Num,
					Title:   &test1Title,
					Body:    &test1Body,
					HTMLURL: &test1URL,
					State:   &test1State,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := api.NewApi(tt.fields.client)
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
		client *http.Client
	}
	client := newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "https://api.github.com/repos/tjgurwara99/Go/pulls" {
			t.Errorf("ListPRs URL = %v, want %v", req.URL, "http://api.github.com/repos/tjgurwara99/Go")
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
	test1Num := 1
	test1Title := "Test Issue 1"
	test1Body := "Test Issue 1 body"
	test1URL := "https://sample.url"
	test1State := "open"
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*github.PullRequest
		wantErr bool
	}{
		{
			name: "Test_api_ListIssues",
			fields: fields{
				client: client,
			},
			args: args{
				repo:  "tjgurwara99/Go",
				state: "open",
			},
			want: []*github.PullRequest{
				{
					Number:  &test1Num,
					Title:   &test1Title,
					Body:    &test1Body,
					HTMLURL: &test1URL,
					State:   &test1State,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := api.NewApi(tt.fields.client)
			got, err := a.ListPRs(tt.args.repo, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPRs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListPRs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_api_GetPR(t *testing.T) {
	cl := newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "https://api.github.com/repos/TheAlgorithms/Go/pulls/1" {
			t.Errorf("GetPR URL = %v, want %v", req.URL, "http://api.github.com/repos/TheAlgorithms/Go/pulls/1")
		}
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
	"number": 1,
	"title": "Test PR 1",
	"body": "Test PR 1 body",
	"html_url": "https://sample.url",
	"state": "open"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	app := api.NewApi(cl)
	got, err := app.GetPR("TheAlgorithms/Go", "1")
	if err != nil {
		t.Errorf("GetPR() error = %v", err)
	}
	prNum := 1
	state := "open"
	prTitle := "Test PR 1"
	prBody := "Test PR 1 body"
	prURL := "https://sample.url"
	want := github.PullRequest{
		Number:  &prNum,
		State:   &state,
		Title:   &prTitle,
		Body:    &prBody,
		HTMLURL: &prURL,
	}
	if *got.HTMLURL != *want.HTMLURL {
		t.Errorf("GetPR HTMLURL don't match; got %v, want %v", *got.HTMLURL, want.HTMLURL)
	}
	if *got.Number != *want.Number {
		t.Errorf("GetPR PR NUMBER don't match; got %v, want %v", *got.HTMLURL, want.HTMLURL)
	}
	if *got.State != *want.State {
		t.Errorf("GetPR PR State don't match; got %v, want %v", *got.State, *want.State)
	}
	if *got.Title != *want.Title {
		t.Errorf("GetPR PR Title don't match; got %v, want %v", *got.Title, *want.Title)
	}
	if *got.Body != *want.Body {
		t.Errorf("GetPR PR Body don't match; got %v, want %v", *got.Body, *want.Body)
	}
}

func Test_api_GetIssue(t *testing.T) {
	cl := newTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != "https://api.github.com/repos/TheAlgorithms/Go/issues/1" {
			t.Errorf("GetIssues URL = %v, want %v", req.URL, "http://api.github.com/repos/TheAlgorithms/Go/issues/1")
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
	app := api.NewApi(cl)
	got, err := app.GetIssue("TheAlgorithms/Go", "1")
	if err != nil {
		t.Errorf("GetIssue() error = %v", err)
	}
	issueNum := 1
	issueTitle := "Test Issue 1"
	issueBody := "Test Issue 1 body"
	issueURL := "https://sample.url"
	issueState := "open"
	want := github.Issue{
		Number:  &issueNum,
		Title:   &issueTitle,
		Body:    &issueBody,
		HTMLURL: &issueURL,
		State:   &issueState,
	}
	if *got.HTMLURL != *want.HTMLURL {
		t.Errorf("GetPR HTMLURL don't match; got %v, want %v", *got.HTMLURL, want.HTMLURL)
	}
	if *got.Number != *want.Number {
		t.Errorf("GetPR PR NUMBER don't match; got %v, want %v", *got.HTMLURL, want.HTMLURL)
	}
	if *got.State != *want.State {
		t.Errorf("GetPR PR State don't match; got %v, want %v", *got.State, *want.State)
	}
	if *got.Title != *want.Title {
		t.Errorf("GetPR PR Title don't match; got %v, want %v", *got.Title, *want.Title)
	}
	if *got.Body != *want.Body {
		t.Errorf("GetPR PR Body don't match; got %v, want %v", *got.Body, *want.Body)
	}
}
