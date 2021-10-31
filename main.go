package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: textify <filename>")
		os.Exit(1)
	}
	if !fileExists(os.Args[1]) {
		fmt.Println("Error: that file does not exist.")
		os.Exit(1)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
