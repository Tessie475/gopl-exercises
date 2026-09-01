// github.go fetches the repository's issues from GitHub. This is the single
// "query GitHub once" the exercise calls for: fetchIssues runs at startup and
// the web handlers then serve everything from memory.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const apiURL = "https://api.github.com"

// User is the subset of a GitHub user we use.
type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

// Milestone is the subset of a milestone we use.
type Milestone struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
}

// Issue is the subset of an issue we use. PullRequest is non-nil when the
// "issue" is actually a pull request, which GitHub returns from this endpoint
// too; we skip those.
type Issue struct {
	Number      int        `json:"number"`
	HTMLURL     string     `json:"html_url"`
	Title       string     `json:"title"`
	State       string     `json:"state"`
	Body        string     `json:"body"`
	User        *User      `json:"user"`
	Milestone   *Milestone `json:"milestone"`
	CreatedAt   time.Time  `json:"created_at"`
	PullRequest *struct{}  `json:"pull_request"`
}

// fetchIssues downloads all issues (open and closed) for a repository,
// following pagination until a short page signals the end.
func fetchIssues(owner, repo string) ([]*Issue, error) {
	var all []*Issue
	for page := 1; ; page++ {
		u := fmt.Sprintf("%s/repos/%s/%s/issues?state=all&per_page=100&page=%d",
			apiURL, owner, repo, page)
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}
		if token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"); token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("github: %s", resp.Status)
		}
		var batch []*Issue
		if err := json.NewDecoder(resp.Body).Decode(&batch); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		all = append(all, batch...)
		if len(batch) < 100 { // a short (or empty) page means we are done
			break
		}
	}
	return all, nil
}
