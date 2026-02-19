package control

import (
	"nyantime/registers"
	"nyantime/util"
	"strconv"
)

func JNE(r *registers.Registers, operator util.Operator, program *[]util.Instruction) error {
	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	if r.GetCompareResult() == false {
		r.SetPC(parsedOperator - 1) // have to set one lower because it increments by one afterwards
	}

	return nil
}
