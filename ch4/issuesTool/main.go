// Command issuesTool creates, reads, updates, and closes GitHub issues from
// the command line (gopl exercise 4.11). It reads a personal access token from
// the GITHUB_PERSONAL_ACCESS_TOKEN environment variable.
//
// Usage:
//
//	issuesTool list   <owner> <repo>
//	issuesTool read   <owner> <repo> <number>
//	issuesTool create <owner> <repo>            (title/body hardcoded for now)
//	issuesTool update <owner> <repo> <number>
//	issuesTool close  <owner> <repo> <number>
//
// The heavy lifting lives in github.go; main just dispatches on the command.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN") == "" {
		log.Fatal("GITHUB_PERSONAL_ACCESS_TOKEN is not set")
	}
	if len(os.Args) < 4 {
		log.Fatal("usage: issuesTool <list|read|create|update|close> <owner> <repo> [number]")
	}
	action, owner, repo := os.Args[1], os.Args[2], os.Args[3]

	switch action {
	case "list":
		issues, err := ListIssues(owner, repo)
		check(err)
		for _, issue := range issues {
			printIssue(issue)
		}

	case "read":
		issue, err := GetIssue(owner, repo, number())
		check(err)
		printIssue(issue)

	case "create":
		// Open an empty editor; first line becomes the title, the rest the body.
		text, err := editText("")
		check(err)
		title, body := splitTitleBody(text)
		if title == "" {
			log.Fatal("aborting: empty title")
		}
		issue, err := CreateIssue(owner, repo, title, body)
		check(err)
		fmt.Printf("created #%d: %s\n", issue.Number, issue.HTMLURL)

	case "update":
		n := number()
		// Fetch the current issue and pre-fill the editor with it, so the
		// user edits in place instead of retyping.
		current, err := GetIssue(owner, repo, n)
		check(err)
		text, err := editText(current.Title + "\n\n" + current.Body)
		check(err)
		title, body := splitTitleBody(text)
		if title == "" {
			log.Fatal("aborting: empty title")
		}
		issue, err := UpdateIssue(owner, repo, n, title, body)
		check(err)
		fmt.Printf("updated #%d: %s\n", issue.Number, issue.HTMLURL)

	case "close":
		issue, err := CloseIssue(owner, repo, number())
		check(err)
		fmt.Printf("closed #%d (state=%s)\n", issue.Number, issue.State)

	default:
		log.Fatalf("unknown action %q (use list, read, create, update, or close)", action)
	}
}

// number reads the issue number from os.Args[4] for the actions that need one.
func number() int {
	if len(os.Args) < 5 {
		log.Fatal("this action needs an issue number: <owner> <repo> <number>")
	}
	n, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatalf("invalid issue number %q: %v", os.Args[4], err)
	}
	return n
}

// check exits with the error printed to stderr, if there is one.
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "issue: %v\n", err)
		os.Exit(1)
	}
}

// printIssue prints an issue's headline and body, guarding the User pointer
// in case it is nil.
func printIssue(issue *Issue) {
	login := ""
	if issue.User != nil {
		login = issue.User.Login
	}
	fmt.Printf("#%d %s [%s] by %s\n%s\n", issue.Number, issue.Title, issue.State, login, issue.Body)
}
