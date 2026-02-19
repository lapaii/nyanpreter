package memory

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func LDI(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
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

	actualValueToLoad, err := util.ParseOperator(secondValueToLoad)

	r.SetAccumulator(actualValueToLoad)

	return nil
}
