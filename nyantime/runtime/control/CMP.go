package control

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

// CMP operator is an address or number
func CMP(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	firstChar := string(operator[0])

	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	switch firstChar {
	case "#", "B", "&":
		// is a user-defined number
		r.SetEqualFlag(r.GetAccumulator() == parsedOperator)
		r.SetGreaterThanFlag(r.GetAccumulator() > parsedOperator)

		return nil
	}

	valueToCompare, err := util.ParseOperator((*program)[parsedOperator].Operator)

	if err != nil {
		return err
	}

	r.SetEqualFlag(r.GetAccumulator() == valueToCompare)
	r.SetGreaterThanFlag(r.GetAccumulator() > valueToCompare)

	return nil
}
