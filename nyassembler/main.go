package main

import (
	"flag"
	"fmt"
	firstpass "nyassembler/first-pass"
	secondpass "nyassembler/second-pass"
	"nyassembler/serialiser"
)

func main() {
	var inputPath string
	var outputPath string
	flag.StringVar(&inputPath, "input", "input.nyan", "the file to assemble")
	flag.StringVar(&outputPath, "output", "output.nyobj", "the output path of the object file")
	flag.Parse()

	StartAssembly(inputPath, outputPath)
}

// Function around the logic so that i can make a
// testing suite that assembles files and tests against their output
func StartAssembly(inputPath string, outputPath string) {
	contents, err := OpenFile(inputPath)

	if err != nil {
		panic(err)
	}

	symbolTable, err := firstpass.FirstPass(contents)

	if err != nil {
		panic(err)
	}

	outputProgram, err := secondpass.SecondPass(contents, symbolTable)

	if err != nil {
		fmt.Println(err)
	}

	bytes := serialiser.Serialise(outputProgram)

	err = WriteFile(outputPath, bytes)

	if err != nil {
		panic(err)
	}
}
