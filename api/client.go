package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type API struct {
	client  *http.Client
	baseUrl string
}

func NewApi(client *http.Client, baseUrl string) *API {
	return &API{
		client:  client,
		baseUrl: baseUrl,
	}
}

type PullRequest struct {
	State  string `json:"state"`
	Body   string `json:"body"`
	Title  string `json:"title"`
	URL    string `json:"html_url"`
	Number int    `json:"number"`
}

type Issue struct {
	State       string      `json:"state"`
	Body        string      `json:"body"`
	Title       string      `json:"title"`
	URL         string      `json:"html_url"`
	Number      int         `json:"number"`
	PullRequest interface{} `json:"pull_request"`
}

func (a *API) GetPR(repo, id string) (*PullRequest, error) {
	url := a.baseUrl + "repos/" + repo + "/pulls/" + id // maybe find a better way to do this
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetPR: error creating http.Request: %w", err)
	}
	res, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetPR: error fetching PR from api.github.com/repos/%s/pulls/%s: %w", repo, id, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request unsuccessful: %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("GetPR: error couldn't read the received body from the api: %w", err)
	}
	var pr PullRequest
	err = json.Unmarshal(body, &pr)
	if err != nil {
		return nil, fmt.Errorf("GetPR: error Unmarshal resp body: %w", err)
	}
	return &pr, nil
}

func (a *API) ListPRs(repo, state string) ([]PullRequest, error) {
	url := a.baseUrl + "repos/" + repo + "/pulls?state=" + state
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("ListPR: error creating http.Request: %w", err)
	}
	res, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error fetching PRs from api.github.com/repos/%s: %w", repo, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request unsuccessful: %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error couldn't read the received body from the api: %w", err)
	}
	var prList []PullRequest
	err = json.Unmarshal(body, &prList)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error Unmarshal resp body: %w", err)
	}
	return prList, nil
}

func (a *API) ListIssues(repo, state string) ([]Issue, error) {
	url := a.baseUrl + "repos/" + repo + "/issues?state=" + state
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error creating http.Request: %w", err)
	}
	res, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error fetching PRs from api.github.com/repos/%s: %w", repo, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request unsuccessful: %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error couldn't read the received body from the api: %w", err)
	}
	var issues []Issue
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, fmt.Errorf("ListPR: error Unmarshal resp body: %w", err)
	}
	var iss []Issue
	for _, issue := range issues {
		if issue.PullRequest != nil {
			continue
		}
		iss = append(iss, issue)
	}
	return iss, nil
}

func (a *API) GetIssue(repo, id string) (*Issue, error) {
	url := a.baseUrl + "repos/" + repo + "/issues/" + id // maybe find a better way to do this
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetPR: error creating http.Request: %w", err)
	}
	res, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetPR: error fetching PR from api.github.com/repos/%s/pulls/%s: %w", repo, id, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request unsuccessful: %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("GetPR: error couldn't read the received body from the api: %w", err)
	}
	var issue Issue
	err = json.Unmarshal(body, &issue)
	if err != nil {
		return nil, fmt.Errorf("GetIssue: error Unmarshal resp body: %w", err)
	}
	return &issue, nil
}
