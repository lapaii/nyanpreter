package shared

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
	CMX
	CMI
	JE
	JNE
	JGE
	JLE
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

var InstructionSet = map[string]Operand{
	"LDM": LDM,
	"LDD": LDD,
	"LDI": LDI,
	"LDX": LDX,
	"LDR": LDR,
	"MOV": MOV,
	"STO": STO,
	"DAT": DAT,
	"ADD": ADD,
	"SUB": SUB,
	"INC": INC,
	"DEC": DEC,
	"JMP": JMP,
	"CMP": CMP,
	"CMX": CMX,
	"CMI": CMI,
	"JE":  JE,
	"JNE": JNE,
	"JGE": JGE,
	"JLE": JLE,
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

// list of instructions that require a user defined number as the operator (#/B/&)
var NumberOperator = []Operand{LDM, LDR, LSL, LSR}

// list of instructions that require addresses
var AddressOperator = []Operand{LDD, LDI, LDX, STO, JMP, CMI, JE, JNE}

// all other instructions can be either defined number or address
