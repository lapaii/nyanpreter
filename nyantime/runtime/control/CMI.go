package control

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func CMI(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	// operator is an address in this instruction

	// parsedOperator is the address to get contents of
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	valueToLoad := (*program)[parsedOperator].Operator

	parsedValue, err := util.ParseOperator(valueToLoad)

	if err != nil {
		return err
	}

	// now lets do it again!
	secondValueToLoad := (*program)[parsedValue].Operator
	actualValueToCompare, err := util.ParseOperator(secondValueToLoad)

	if r.GetAccumulator() == actualValueToCompare {
		r.SetCompareResult(true)
	} else {
		r.SetCompareResult(false)
	}

	return nil
}
