package main

import (
	"shared"
)

func reverseMap(m map[string]shared.Opcode) map[shared.Opcode]string {
	n := make(map[shared.Opcode]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}
