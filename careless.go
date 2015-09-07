package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/jzelinskie/geddit"
)

const (
	TRIGGER = "I could care less"
	SLEEP_TIME = time.Second * 5
)

var checked map[string]struct{}

func main () {
	checked = make(map[string]struct{})
	var subreddits = [...]string {"askreddit", "4chan"} // subreddits to check

	// login to reddit
	r, err := geddit.NewLoginSession (USERNAME, PASSWORD, "CarelessBot v1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to log in\n")
		return
	}

	// set listing options
	subOpts := geddit.ListingOptions{
		Limit: 100,
	}

	for {
		for _, sub := range subreddits {
			submissions, err := r.SubredditSubmissions(sub, geddit.NewSubmissions, subOpts)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get new submissions for %s\n", sub)
			} else {
				for _, s := range submissions {
					fmt.Println("Checking comments in ", s.Permalink)
					// get submission comments
					comments, err := r.Comments(s)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Failed to get submission comments: %s\n", err.Error())
					} else {
						go checkComments(comments)
					}
				}
			}
		}
		// sleep for a while
		fmt.Println("Sleeping for", SLEEP_TIME)
		time.Sleep(SLEEP_TIME)
	}
}

func checkComments (comments []*geddit.Comment) {
	for _, c := range comments {
		// check if we have already checked this comment
		_, ok := checked[c.FullID]
		if !ok {
			if strings.Contains(c.Body, TRIGGER) {
				fmt.Println("Found an offender: ", c.FullID)
			}
			checked[c.FullID] = struct{}{}
		} else {
			fmt.Println("Skipping comment already scanned: ", c.FullID)
		}
	}
}
