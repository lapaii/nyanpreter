package math

import (
	"nyantime/registers"
	"shared"
)

// operator is IDX or ACC
func DEC(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator := string(operator)

	if parsedOperator == "IDX" {
		r.DecrementIndex(1)
	}

	if parsedOperator == "ACC" {
		r.DecrementAccumulator(1)
	}

	return nil
}
