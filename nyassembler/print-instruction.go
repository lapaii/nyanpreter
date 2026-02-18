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

func PrintInstruction(inst secondpass.Instruction) {
	reversedMap := reverseMap(secondpass.InstructionSet)

	fmt.Printf("%s %s\n", reversedMap[inst.Operand], inst.Operator)
}
