package decoder

import (
	"nyantime/util"
	"shared"
)

func DecodeInstructions(program []byte) []shared.Instruction {
	instructions := util.SplitSlice(program, byte(0))

	outputProgram := []shared.Instruction{}

	for _, inst := range instructions {
		operand := inst[0]
		operator := inst[1:]

		outputProgram = append(outputProgram, shared.Instruction{
			Operand:  shared.Operand(operand),
			Operator: shared.Operator(operator),
		})
	}

	return outputProgram
}
