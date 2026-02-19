package memory

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func LDD(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	// operator is an address in this instruction

	// parsedOperator is the address to load from
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	valueToLoad := (*program)[parsedOperator].Operator

	parsedValue, err := util.ParseOperator(valueToLoad)

	if err != nil {
		return err
	}

	r.SetAccumulator(parsedValue)

	return nil
}
