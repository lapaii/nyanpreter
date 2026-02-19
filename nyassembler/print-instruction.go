package main

import (
	"fmt"
	"shared"
)

func reverseMap(m map[string]shared.Operand) map[shared.Operand]string {
	n := make(map[shared.Operand]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func PrintInstruction(inst shared.Instruction, idx int) {
	reversedMap := reverseMap(shared.InstructionSet)

	fmt.Printf("%d: %s %s\n", idx, reversedMap[inst.Operand], inst.Operator)
}
