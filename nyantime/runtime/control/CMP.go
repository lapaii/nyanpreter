package control

import (
	"nyantime/registers"
	"nyantime/util"
	"shared"
)

// CMP operator is an address
func CMP(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := util.ParseOperator(operator) // just line number

	valueToCompare, err := util.ParseOperator((*program)[parsedOperator].Operator)

	if err != nil {
		return err
	}

	if r.GetAccumulator() == valueToCompare {
		r.SetCompareResult(true)
	} else {
		r.SetCompareResult(false)
	}

	return nil
}
