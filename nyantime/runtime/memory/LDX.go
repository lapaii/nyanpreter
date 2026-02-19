package memory

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func LDX(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	// this is an address
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	addressToLoad := parsedOperator + r.GetIndex()

	valueToSet, err := util.ParseOperator((*program)[addressToLoad].Operator)

	if err != nil {
		return err
	}

	r.SetAccumulator(valueToSet)

	return nil
}
