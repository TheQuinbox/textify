// All our parsing code. Each function takes the text as a string, and returns a string.

package main

import (
	"jaytaylor.com/html2text"
	"github.com/writeas/go-strip-markdown"
	"fmt"
)

func parseMarkdown(text string) string {
	data := stripmd.Strip(string(text))
	return data
}

func parseHtml(text string) string {
	data, err := html2text.FromString(string(text), html2text.Options{PrettyTables: true})
	if err != nil {
		fmt.Printf("Error writing the file. %s\n", err)
		return ""
	}
	return data
}
