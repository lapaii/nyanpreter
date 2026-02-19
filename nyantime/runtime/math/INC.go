package math

import (
	"nyantime/registers"
	"shared"
)

// operator is IDX or ACC
func INC(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	parsedOperator := string(operator)

	if parsedOperator == "IDX" {
		r.IncrementIndex(1)
	}

	if parsedOperator == "ACC" {
		r.IncrementAccumulator(1)
	}

	return nil
}
