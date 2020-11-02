package cli

import (
	"errors"
	"flag"
	"strings"
)

// Cli returning tag array, url and error
func Cli() ([]string, string, error) {
	var tagFlag = flag.String("tags", "", "add html tags, example usage --tags 'p span a' ")
	var urlFlag = flag.String("url", "", "add a url, example usage --url 'http://example.com' ")

	flag.Parse()
	var tags string = *tagFlag
	var url string = *urlFlag

	if url == "" || tags == "" {
		return nil, "", errors.New("you have entered wrong tag or url, please try again")
	}

	textTags := strings.Fields(tags)

	return textTags, url, nil
}
