package parser

import (
	"fmt"
	"nyanpreter/instructions"
	"slices"
)

// this doesn't check if the operator is valid for use
func ParseOperator(lineNum int, instruction instructions.Operand, part string) (string, error) {
	// if the instruction doesnt use an operator
	if slices.Contains(instructions.NoOperator, instruction) {
		return "", nil
	} else if part == "" {
		return "", fmt.Errorf("instruction %d on line %d requires an operator!\n", instruction, lineNum)
	}

	// check if the operator is of the right type for the instruction
	if slices.Contains(instructions.RegisterOperator, instruction) {
		// instruction needs a register operator
		// is the operator either one of them?
		if part == "ACC" || part == "IDX" {
			// yes it is!!
			return part, nil
		}

		return "", fmt.Errorf(
			"instruction %d on line %d requires a register as the operator, instead got %s",
			instruction, lineNum, part,
		)
	}

	return part, nil
}
