package util

type Operand uint8
type Operator string

type Instruction struct {
	Operand  Operand
	Operator Operator
}

func ModifyInstruction(instructions *[]Instruction, index int, newInstruction Instruction) {
	(*instructions)[index] = newInstruction
}

func (inst *Instruction) ModifyOperator(newOperator string) {
	inst.Operator = Operator(newOperator)
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
	JE
	JNE
	IN
	OUT
	END

	// bitwise operations
	AND
	XOR
	OR
	LSL
	LSR
)
