package instructions

type Operand int

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

var InstructionSet = map[string]Operand{
	"LDM": LDM,
	"LDD": LDD,
	"LDI": LDI,
	"LDX": LDX,
	"LDR": LDR,
	"MOV": STO,
	"ADD": ADD,
	"SUB": SUB,
	"INC": INC,
	"DEC": DEC,
	"JMP": JMP,
	"CMP": CMP,
	"CMI": CMI,
	"JPE": JPE,
	"JPN": JPN,
	"IN":  IN,
	"OUT": OUT,
	"END": END,
	"AND": AND,
	"XOR": XOR,
	"OR":  OR,
	"LSL": LSL,
	"LSR": LSR,
}

// list of instructions that dont use an operator
var NoOperator = []Operand{IN, OUT, END}

// list of instructions which the operator needs to be a register (ACC or IDX)
var RegisterOperator = []Operand{MOV, INC, DEC}
