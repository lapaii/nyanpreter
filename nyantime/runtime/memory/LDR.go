package memory

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

func LDR(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator)

	if err != nil {
		return err
	}

	r.SetIndex(parsedOperator)

	return nil
}
