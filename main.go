package main

import (
	"flag"
	"nyanpreter/parser"
)

func main() {
	// gets the input file from flags
	var filePath string
	flag.StringVar(&filePath, "input", "input.nyan", "the source file to run")

	flag.Parse()

	fileContents, err := parser.LoadFile(filePath)

	if err != nil {
		panic(err)
	}

	_, err = parser.Parse(fileContents)

	if err != nil {
		panic(err)
	}
}
