package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("textify", "Converts many common document types to plane text")
	input := parser.String("i", "input", &argparse.Options{Required: true, Help: "File to convert"})
	output := parser.String("o", "output", &argparse.Options{Required: false, Help: "File to write the output to. Default is input-file + .txt"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
	}
	fmt.Println("Got input file of: " + *input + " and output file of " + *output)
}
