package memory

import (
	"nyantime/registers"
	"nyantime/util"
)

func LDM(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator, *program)

	if err != nil {
		return err
	}

	r.SetAccumulator(parsedOperator)

	return nil
}
