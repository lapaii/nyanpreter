package runtime

import (
	"fmt"
	"nyantime/registers"
	"nyantime/util"
)

func Runtime(decodedProgram []util.Instruction) error {
	registers := registers.Registers{
		ProgramCounter: 0,
		Accumulator:    0,
		Index:          0,
	}

	shouldContinue := true

	for shouldContinue {
		currentInstruction := decodedProgram[registers.ProgramCounter]

		if currentInstruction.Operand == util.END {
			shouldContinue = false
			continue
		}

		err := FunctionMap[currentInstruction.Operand](&registers, currentInstruction.Operator, decodedProgram)

		if err != nil {
			return err
		}

		fmt.Println(registers.Accumulator)

		registers.IncrementPC()
	}

	return nil
}
