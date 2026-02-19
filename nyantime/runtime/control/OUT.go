package control

import (
	"fmt"
	"nyantime/registers"
	"nyantime/util"
)

func OUT(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	fmt.Printf("%c", r.Accumulator)

	return nil
}
