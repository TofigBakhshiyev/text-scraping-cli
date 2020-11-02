package tokenize

import (
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Tokenize takes tag array and url as a input, and returns text of html
func Tokenize(inputTextTags []string, url string) []string {
	textTags := inputTextTags
	var text []string

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	tag := ""
	attrs := map[string]string{}
	enter := false

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false
			attrs = map[string]string{}

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					for _, attr := range token.Attr {
						attrs[attr.Key] = attr.Val
					}
					break
				}
			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					//fmt.Printf("html tag: %v text: %v\n", tag, data)
					text = append(text, data)
				}
			}
		}
	}

	return text
}
