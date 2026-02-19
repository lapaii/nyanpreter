package bitwise

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

// LSR can only be a defined number
func LSR(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	currentACC := r.GetAccumulator()

	r.SetAccumulator(currentACC >> parsedOperator)

	return nil
}
