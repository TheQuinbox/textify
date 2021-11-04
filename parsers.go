// All our parsing code. Each function takes the filename as a string, and returns a string containing the parsed text.

package main

import (
	"github.com/writeas/go-strip-markdown"
	"jaytaylor.com/html2text"
)

func parseMarkdown(text string) string {
	data := stripmd.Strip(string(text))
	return data
}

func parseHtml(text string) string {
	data, err := html2text.FromString(string(text), html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	return data
}
