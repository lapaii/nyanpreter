package control

import (
	"nyantime/registers"
	"shared"
	"strconv"
)

func JMP(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	r.SetPC(parsedOperator - 1) // 1 below so that pc increments to right place

	return nil
}
