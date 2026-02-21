package operand

import (
	"fmt"
	firstpass "nyassembler/first-pass"
	"shared"
	"strconv"
)

func ParseOperand(idx int, operand string, symbolTable firstpass.SymbolTable) (shared.Operand, shared.OperandType, error) {
	operandLen := len(operand) - 1
	dereference := false

	// dereference
	if operand[0] == '(' && operand[operandLen] == ')' {
		operand = operand[1:operandLen]
		dereference = true
	}

	// register
	if operand[0] == '%' {
		registerNum := shared.RegisterMap[operand]

		// register doesnt exist!
		if registerNum == 0 {
			return 0, 0, fmt.Errorf("[line %d] register %s doesn't exist!", idx+1, operand)
		}

		operandType := shared.Register
		if dereference {
			operandType = shared.RegisterPointer
		}

		// return, subtracting 1 from registerNum so its correct
		return shared.Operand(registerNum - 1), operandType, nil
	}

	// immediate
	if operand[0] == '$' {
		operand = operand[1:]
		var operandValue int64

		numBase := 10

		if len(operand) >= 2 && operand[0:2] == "0x" {
			numBase = 16
			operand = operand[2:]
		}

		var err error
		operandValue, err = strconv.ParseInt(operand, numBase, 0)

		if err != nil {
			return 0, 0, fmt.Errorf("[line %d] %s is an invalid number!", idx+1, operand)
		}

		operandType := shared.Immediate
		if dereference {
			operandType = shared.ImmediatePointer
		}

		return shared.Operand(operandValue), operandType, nil
	}

	// label
	lineEntry := symbolTable[operand]

	// line entry 0 != symbol does exist
	if lineEntry != 0 {
		// correct it to the actual line
		lineEntry = lineEntry - 1

		operandType := shared.Immediate
		if dereference {
			operandType = shared.ImmediatePointer
		}

		return shared.Operand(lineEntry), operandType, nil
	}

	return 0, 0, fmt.Errorf("[line %d] operand %s is used like a label but isnt defined!", idx+1, operand)
}
