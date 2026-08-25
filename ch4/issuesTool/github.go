// github.go holds the GitHub client: the issue types, two shared helpers
// (newRequest builds a request with auth + headers; do sends it, checks the
// status, and decodes the reply), and one function per CRUD operation.
//
// Every operation is the same five parts (method, URL, body, status, decode),
// so the helpers carry the repetition and each function just supplies what
// differs.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const apiURL = "https://api.github.com"

// User is the subset of a GitHub user we use.
type User struct {
	Login string `json:"login"`
}

// Issue is the subset of an issue we use. The json tags match GitHub's keys.
type Issue struct {
	Number    int       `json:"number"`
	HTMLURL   string    `json:"html_url"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Body      string    `json:"body"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

// issuesURL is the collection URL: .../repos/{owner}/{repo}/issues
func issuesURL(owner, repo string) string {
	return fmt.Sprintf("%s/repos/%s/%s/issues", apiURL, owner, repo)
}

// issueURL points at one issue: .../repos/{owner}/{repo}/issues/{number}
func issueURL(owner, repo string, number int) string {
	return fmt.Sprintf("%s/repos/%s/%s/issues/%d", apiURL, owner, repo, number)
}

// newRequest builds a request carrying GitHub's headers. When body is non-nil
// it is marshalled to JSON and a Content-Type header is added.
func newRequest(method, url string, body any) (*http.Request, error) {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(data)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	if token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"); token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// do sends req, requires the status code to equal want, and (when out is
// non-nil) decodes the JSON body into it. On a wrong status it returns an
// error carrying GitHub's response text, which usually explains the problem.
func do(req *http.Request, want int, out any) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != want {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s: %s", resp.Status, body)
	}
	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}

// CreateIssue opens a new issue. (POST, expects 201 Created.)
func CreateIssue(owner, repo, title, body string) (*Issue, error) {
	req, err := newRequest("POST", issuesURL(owner, repo),
		map[string]string{"title": title, "body": body})
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err := do(req, http.StatusCreated, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// GetIssue reads one issue by number. (GET, expects 200 OK.)
func GetIssue(owner, repo string, number int) (*Issue, error) {
	req, err := newRequest("GET", issueURL(owner, repo, number), nil)
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err := do(req, http.StatusOK, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// ListIssues returns a repository's issues. (GET, expects 200 OK.)
func ListIssues(owner, repo string) ([]*Issue, error) {
	req, err := newRequest("GET", issuesURL(owner, repo), nil)
	if err != nil {
		return nil, err
	}
	var issues []*Issue
	if err := do(req, http.StatusOK, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

// UpdateIssue changes an issue's title and body. (PATCH, expects 200 OK.)
func UpdateIssue(owner, repo string, number int, title, body string) (*Issue, error) {
	req, err := newRequest("PATCH", issueURL(owner, repo, number),
		map[string]string{"title": title, "body": body})
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err := do(req, http.StatusOK, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// CloseIssue closes an issue. GitHub's REST API has no delete-issue endpoint,
// so closing (state = "closed") is the equivalent. (PATCH, expects 200 OK.)
func CloseIssue(owner, repo string, number int) (*Issue, error) {
	req, err := newRequest("PATCH", issueURL(owner, repo, number),
		map[string]string{"state": "closed"})
	if err != nil {
		return nil, err
	}
	var issue Issue
	if err := do(req, http.StatusOK, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
