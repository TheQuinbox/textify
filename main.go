package main

import (
	"fmt"
	"github.com/writeas/go-strip-markdown"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}
	if !fileExists(os.Args[1]) {
		fmt.Printf("Error: %s does not exist.\n", os.Args[1])
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading the file. %s\n", err)
		os.Exit(1)
	}
	if !strings.HasSuffix(strings.ToLower(os.Args[1]), ".md") {
		fmt.Println("Error: textify doesn't support that type of file.")
		os.Exit(1)
	}
	fmt.Printf("Converting %s to text...", os.Args[1])
	err = os.WriteFile(fmt.Sprintf("%s.txt", os.Args[1]), []byte(stripmd.Strip(string(data))), 0644)
	if err != nil {
		fmt.Printf("There was an error writing the file. %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}

// Simple function to check if a file exists. If the file doesn't exist or is a directory, it returns false. Otherwise, true.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
