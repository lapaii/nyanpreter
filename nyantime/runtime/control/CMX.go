package control

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

// CMX operator is an address
func CMX(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator) // just line number

	valueToCompare, err := util.ParseOperator((*program)[parsedOperator].Operator)

	if err != nil {
		return err
	}

	if r.GetIndex() == valueToCompare {
		r.SetCompareResult(true)
	} else {
		r.SetCompareResult(false)
	}

	return nil
}
