package secondpass

import "fmt"

func ParseOperand(input string, lineNum int) (Operand, error) {
	var operand Operand
	var returnError error = nil

	operand = InstructionSet[input]

	if operand == INVALID {
		returnError = fmt.Errorf("Invalid instruction at Line %d: %s", lineNum+1, input)
	}

	return operand, returnError
}
