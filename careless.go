package main

import (
	"fmt"
	"os"
	"github.com/jzelinskie/geddit"
)

func main () {
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

	for _, sub := range subreddits {
		submissions, err := r.SubredditSubmissions(sub, geddit.NewSubmissions, subOpts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get new submissions for %s\n", sub)
		} else {
			for _, s := range submissions {
				fmt.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
			}
		}
	}
}
