// Exercise 4.10: modify issues to report results in age categories: less than
// a month old, less than a year old, and more than a year old.
//
// The categories are mutually exclusive: a switch stops at the first matching
// case, so each issue is counted in exactly one bucket and the three totals
// sum to result.TotalCount.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// Compute the boundaries once, so every comparison uses the same "now".
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)

	lessThanAMonthAgo := 0
	lessThanAYearAgo := 0 // older than a month but younger than a year
	moreThanAYearAgo := 0
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(monthAgo):
			lessThanAMonthAgo++
		case item.CreatedAt.After(yearAgo):
			lessThanAYearAgo++
		default:
			moreThanAYearAgo++
		}
	}
	fmt.Printf("Issues opened less than a month ago: %d\n", lessThanAMonthAgo)
	fmt.Printf("Issues opened less than a year ago: %d\n", lessThanAYearAgo)
	fmt.Printf("Issues opened more than a year ago: %d\n", moreThanAYearAgo)
}
