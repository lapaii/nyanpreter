package memory

import (
	"fmt"
	"nyantime/registers"
	"shared"
	"strconv"
)

func STO(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	// sto just sets the value of the operator on the instruction
	// that this instructions operator points to the value of the acc

	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	shared.ModifyInstruction(program, parsedOperator, shared.Instruction{
		Operand:  (*program)[parsedOperator].Operand,
		Operator: shared.Operator(fmt.Sprintf("#%d", r.GetAccumulator())),
	})

	return nil
}
