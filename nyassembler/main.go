package main

import (
	"flag"
	"fmt"
	firstpass "nyassembler/first-pass"
	secondpass "nyassembler/second-pass"
)

func main() {
	var inputPath string
	var outputPath string
	flag.StringVar(&inputPath, "input", "input.nyan", "the file to assemble")
	flag.StringVar(&outputPath, "output", "output.nyobj", "the output path of the object file")

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

	fmt.Println(outputProgram)

	for idx, inst := range outputProgram {
		PrintInstruction(inst, idx)
	}
}
