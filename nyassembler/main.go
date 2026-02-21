package main

import (
	"flag"
	firstpass "nyassembler/first-pass"
	secondpass "nyassembler/second-pass"
	"nyassembler/serialiser"
	"os"
)

func main() {
	var inputPath string
	var outputPath string
	flag.StringVar(&inputPath, "input", "input.nyasm", "the file to assemble")
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
		panic(err)
	}

	buf, err := serialiser.Serialise(outputProgram)

	if err != nil {
		panic(err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf.WriteTo(file)
}
