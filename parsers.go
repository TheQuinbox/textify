// All our parsing code. Each function takes the filename as a string, and returns a string containing the parsed text.

package main

import (
	"github.com/writeas/go-strip-markdown"
	"jaytaylor.com/html2text"
	"os"
)

func parseMarkdown(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	final := stripmd.Strip(string(data))
	return final
}

func parseHtml(filename string) string {
	text, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data, err := html2text.FromString(string(text), html2text.Options{PrettyTables: true, TextOnly: true})
	if err != nil {
		panic(err)
	}
	return data
}
