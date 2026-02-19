package control

import (
	"fmt"
	"nyantime/registers"
	"shared"
)

func OUT(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	fmt.Printf("%c", r.GetAccumulator())

	return nil
}
