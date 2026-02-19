package secondpass

import (
	"fmt"
	"shared"
)

func ParseOperand(input string, lineNum int) (shared.Operand, error) {
	var operand shared.Operand
	var returnError error = nil

	operand = shared.InstructionSet[input]

	if operand == shared.INVALID {
		returnError = fmt.Errorf("Invalid instruction at Line %d: %s", lineNum+1, input)
	}

	return operand, returnError
}
