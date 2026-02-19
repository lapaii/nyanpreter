package math

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

// add can be a number or address
func ADD(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	firstChar := string(operator[0])

	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	switch firstChar {
	case "#", "B", "&":
		// is a user-defined number
		r.IncrementAccumulator(parsedOperator)

		return nil
	}

	valueToAdd, err := util.ParseOperator((*program)[parsedOperator].Operator)

	if err != nil {
		return err
	}

	r.IncrementAccumulator(valueToAdd)

	return nil
}
