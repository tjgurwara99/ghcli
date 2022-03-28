package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
)

type API struct {
	client *http.Client
}

func NewApi(client *http.Client) *API {
	return &API{
		client: client,
	}
}

func (a *API) GetPR(repo, id string) (*github.PullRequest, error) {
	client := github.NewClient(a.client)
	owner, repo, err := getOwnerAndRepo(repo)
	if err != nil {
		return nil, fmt.Errorf("GetPR: %w", err)
	}
	prID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("GetPR: id must be an integer: %w", err)
	}
	pr, resp, err := client.PullRequests.Get(context.Background(), owner, repo, prID)
	if err != nil {
		return nil, fmt.Errorf("GetPR: retrieving PR: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GetPR: non 200 response: %s", resp.Status)
	}
	return pr, nil
}

func (a *API) GetIssue(repo, id string) (*github.Issue, error) {
	client := github.NewClient(a.client)
	owner, repo, err := getOwnerAndRepo(repo)
	if err != nil {
		return nil, fmt.Errorf("GetIssue: %w", err)
	}
	issueID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("GetIssue: id must be an integer: %w", err)
	}
	issue, resp, err := client.Issues.Get(context.Background(), owner, repo, issueID)
	if err != nil {
		return nil, fmt.Errorf("GetIssue: retrieving PR: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GetIssue: non 200 response: %s", resp.Status)
	}
	return issue, nil

}

func getOwnerAndRepo(addr string) (owner string, repo string, err error) {
	split := strings.Split(addr, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("incorrect input format - repo should be provided along with owner eg 'owner/repo'")
	}
	return split[0], split[1], nil
}

func (a *API) ListPRs(repo, state string) ([]*github.PullRequest, error) {
	client := github.NewClient(a.client)
	owner, repo, err := getOwnerAndRepo(repo)
	if err != nil {
		return nil, fmt.Errorf("ListPRS: %w", err)
	}
	prs, resp, err := client.PullRequests.List(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, fmt.Errorf("ListPRs: error retrieving PRs: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ListPRs: non successful response code: %s", resp.Status)
	}
	return prs, nil
}

func (a *API) ListIssues(repo, state string) ([]*github.Issue, error) {
	client := github.NewClient(a.client)
	owner, repo, err := getOwnerAndRepo(repo)
	if err != nil {
		return nil, fmt.Errorf("ListIssues: %w", err)
	}
	opt := github.IssueListByRepoOptions{
		State: state,
	}
	issues, resp, err := client.Issues.ListByRepo(context.TODO(), owner, repo, &opt)
	if err != nil {
		return nil, fmt.Errorf("ListIssues: error retrieving issues: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ListIssues: non successful response code: %s", resp.Status)
	}
	return issues, nil
}
