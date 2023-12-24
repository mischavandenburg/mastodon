package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mischavandenburg/mastodon"
	"github.com/mischavandenburg/mastodon/internal/feed"
)

var tags string

// gets the latest post from RSS feed and posts it to Mastodon
func main() {

	// check if tags are given and format them
	if len(os.Args) > 1 {
		tags = "\n" + strings.Join(os.Args[1:], " ")
	}

	// retrieve the latest post from my blog
	r := feed.GetLatestPost()

	// format the post to Mastodon
	toot := r + tags

	// The access token to authenticate with Mastodon stored in an env variable
	token := os.Getenv("MASTODON_TOKEN")

	// Call the postToMastodon function and print the result or error
	result, err := mastodon.PostToMastodon(toot, token)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success:", result)
	}
}
