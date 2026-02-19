package decoder

import (
	"nyantime/util"
)

func DecodeInstructions(program []byte) []util.Instruction {
	instructions := util.SplitSlice(program, byte(0))

	outputProgram := []util.Instruction{}

	for _, inst := range instructions {
		operand := inst[0]
		operator := inst[1:]

		outputProgram = append(outputProgram, util.Instruction{
			Operand:  util.Operand(operand),
			Operator: util.Operator(operator),
		})
	}

	return outputProgram
}
