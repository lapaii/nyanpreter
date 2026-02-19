package runtime

import (
	"fmt"
	"nyantime/registers"
	"nyantime/util"
)

func Runtime(decodedProgram []util.Instruction) error {
	registers := registers.NewRegisters()
	shouldContinue := true

	for shouldContinue {
		currentInstruction := decodedProgram[registers.GetPC()]

		if currentInstruction.Operand == util.END {
			shouldContinue = false
			continue
		}

		err := FunctionMap[currentInstruction.Operand](&registers, currentInstruction.Operator, &decodedProgram)

		if err != nil {
			return fmt.Errorf("failed on instruction: %+v\nerror: %s", currentInstruction, err.Error())
		}

		registers.IncrementPC()
	}

	return nil
}
