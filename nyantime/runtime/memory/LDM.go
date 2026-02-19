package memory

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func LDM(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	r.SetAccumulator(parsedOperator)

	return nil
}
