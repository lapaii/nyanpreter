package serialiser

import (
	secondpass "nyassembler/second-pass"
)

func Serialise(input []secondpass.Instruction) []byte {
	// this is whats going to output a buffer with the binary format

	// 1 byte for operand
	// as many bytes as needed for operator, followed by a null byte

	buf := []byte{}

	for _, instruction := range input {
		// write byte for operand
		buf = append(buf, byte(instruction.Operand))

		// write bytes for operator
		buf = append(buf, []byte(instruction.Operator)...)

		buf = append(buf, 0)
	}

	return buf
}
