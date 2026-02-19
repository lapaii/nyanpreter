package control

import (
	"nyantime/registers"
	"shared"
	"strconv"
)

func JNE(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	if r.GetCompareResult() == false {
		r.SetPC(parsedOperator - 1) // have to set one lower because it increments by one afterwards
	}

	return nil
}
