package bitwise

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func AND(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	firstChar := string(operator[0])

	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	currentACC := r.GetAccumulator()

	switch firstChar {
	case "#", "B", "&":
		// is a user-defined number
		r.SetAccumulator(currentACC & parsedOperator)

		return nil
	}

	valueToAnd, err := util.ParseOperator((*program)[parsedOperator].Operator)

	r.SetAccumulator(currentACC & valueToAnd)

	return nil
}
