package parser

import (
	"nyanpreter/instructions"
	"regexp"
	"strings"
)

type Instruction struct {
	Operand  instructions.Operand
	Operator string
}

func Parse(contents []string) ([]Instruction, error) {
	// go through each line
	// clean any comments, `//` is comment sign
	// split into operand / operator
	// check if operand is legit and convert it into the enum type
	// check if operator exists for the instructions that need it
	// return instruction array

	// regex to match all characters after a // for removing comments
	commentRegex := regexp.MustCompile(`(\/\/.*)`)

	parsedInstructions := []Instruction{}

	for idx, line := range contents {
		// remove comments
		cleansedLine := commentRegex.ReplaceAllString(line, "")

		parts := strings.Split(cleansedLine, " ")

		// parse operand
		parsedOperand, err := ParseOperand(idx, parts[0])

		if err != nil {
			return []Instruction{}, err
		}

		// parse operator
		parsedOperator, err := ParseOperator(idx, parsedOperand, parts[1])

		if err != nil {
			return []Instruction{}, err
		}

		// add parsed instruction to slice
		parsedInstructions = append(parsedInstructions, Instruction{
			Operand:  parsedOperand,
			Operator: parsedOperator,
		})
	}

	return parsedInstructions, nil
}
