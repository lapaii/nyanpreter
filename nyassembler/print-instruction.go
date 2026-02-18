package main

import (
	"fmt"
	secondpass "nyassembler/second-pass"
)

func reverseMap(m map[string]secondpass.Operand) map[secondpass.Operand]string {
	n := make(map[secondpass.Operand]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func PrintInstruction(inst secondpass.Instruction, idx int) {
	reversedMap := reverseMap(secondpass.InstructionSet)

	fmt.Printf("%d: %s %s\n", idx, reversedMap[inst.Operand], inst.Operator)
}
