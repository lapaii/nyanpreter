package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Operand  string
	Operator string
}

func main() {
	// gets the input file from flags
	var filePath string
	flag.StringVar(&filePath, "input", "input.nyan", "the source file to run")

	flag.Parse()

	// read the file and get it ready for parsing
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	content := string(contentBytes)

	lines := strings.Split(content, "\n")

	// file now ready for some "parsing"
	instructions := []Instruction{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		operand := CheckPart(parts[0])
		operator := ""

		// check if operand is one that doesnt have an operator
		if operand != "OUT" {
			operator = CheckPart(parts[1])
		}

		instructions = append(instructions, Instruction{
			Operand:  operand,
			Operator: operator,
		})
	}

	// file is parsed
	// time to "interpret" it

	accumulator := 0

	for _, instruction := range instructions {
		// immediate addressing, LDM #n where n is loaded into ACC
		if instruction.Operand == "LDM" {
			value, err := strconv.Atoi(strings.ReplaceAll(instruction.Operator, "#", ""))

			if err != nil {
				panic(err)
			}

			accumulator = value
		}

		// add instruction, ADD #n/Bn/&n to ACC
		if instruction.Operand == "ADD" {
			// base 10 number
			if instruction.Operator[0] == byte('#') {
				value, err := strconv.Atoi(instruction.Operator[1:])

				if err != nil {
					panic(err)
				}

				accumulator = accumulator + value
			}
		}

		// output whatever is in the accumulator to stdout
		if instruction.Operand == "OUT" {
			fmt.Printf("%c\n", accumulator)
		}
	}
}

func CheckPart(part string) string {
	if part == "//" {
		return ""
	}

	return part
}
