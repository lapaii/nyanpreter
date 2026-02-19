package memory

import (
	"fmt"
	"nyantime/registers"
	"nyantime/util"
)

func LDM(r *registers.Registers, operator util.Operator, program []util.Instruction) error {
	parsedOperator := util.ParseOperator(operator)

	fmt.Println(parsedOperator)

	return nil
}
