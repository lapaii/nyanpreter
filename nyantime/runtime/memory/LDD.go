package memory

import (
	"nyantime/registers"
	"nyantime/util"
	"strconv"
)

func LDD(r *registers.Registers, operator util.Operator, program []util.Instruction) error {
	// operator is an address in this instruction so
	// i know its just a line number i can index to

	parsedOperator, err := strconv.Atoi(string(operator))

	if err != nil {
		return err
	}

	valueToLoad := program[parsedOperator].Operator

	parsedValue := util.ParseOperator(valueToLoad)

	r.SetAccumulator(parsedValue)

	return nil
}
