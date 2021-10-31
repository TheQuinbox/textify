package main

import (
	"fmt"
	"github.com/writeas/go-strip-markdown"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: textify <filename>")
		os.Exit(1)
	}
	if !fileExists(os.Args[1]) {
		fmt.Println("Error: that file does not exist.")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(stripmd.Strip(string(data)))
}

// Simple function to check if a file exists. If the file doesn't exist or is a directory, it returns false. Otherwise, true.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
