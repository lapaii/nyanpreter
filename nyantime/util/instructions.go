package util

type Operand uint8
type Operator string

type Instruction struct {
	Operand  Operand
	Operator Operator
}

const (
	INVALID Operand = iota

	// memory instructions
	LDM
	LDD
	LDI
	LDX
	LDR
	MOV
	STO

	DAT

	// maths
	ADD
	SUB
	INC
	DEC

	// control flow
	JMP
	CMP
	CMI
	JPE
	JPN
	IN
	OUT
	END

	// logic operations
	AND
	XOR
	OR
	LSL
	LSR
)
