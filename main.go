package main

import (
	"fmt"
	"log"
	"textscraping/cli"
	"textscraping/tokenize"
)

func main() {
	// getting tags and url from command line
	textTags, url, err := cli.Cli()

	if err != nil {
		log.Fatal(err)
	}

	// getting text data
	text := tokenize.Tokenize(textTags, url)

	fmt.Println(text)
}
