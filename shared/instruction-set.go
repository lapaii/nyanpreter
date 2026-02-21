package shared

type Opcode uint8
type Operand int32

type OperandType uint8

const (
	Register         OperandType = iota // 00
	RegisterPointer                     // 01
	Immediate                           // 10
	ImmediatePointer                    // 11
)

type Instruction struct {
	SourceType OperandType
	DestType   OperandType
	Opcode     Opcode
	Source     Operand
	Dest       Operand
}

func ModifyInstruction(instructions *[]Instruction, index int, newInstruction Instruction) {
	(*instructions)[index] = newInstruction
}

func (inst *Instruction) ModifyOperator(newOperator string) {
	// inst.Source = Operator(newOperator)
}

const (
	INVALID Opcode = iota

	mov

	add
	sub
	mul
	div
	mod

	and
	not
	xor
	or
	shr
	shl

	jmp
	cmp
	je
	jne
	jz
	jnz
	jge
	jg
	jle
	jl

	in
	out
	outd
	end
	call
	ret

	dn
)

type OpcodeInfo struct {
	Opcode     Opcode
	NumOperand uint8
}

var InstructionSet = map[string]OpcodeInfo{
	"mov": {
		Opcode:     mov,
		NumOperand: 2,
	},
	"add": {
		Opcode:     add,
		NumOperand: 2,
	},
	"sub": {
		Opcode:     sub,
		NumOperand: 2,
	},
	"mul": {
		Opcode:     mul,
		NumOperand: 2,
	},
	"div": {
		Opcode:     div,
		NumOperand: 2,
	},
	"mod": {
		Opcode:     mod,
		NumOperand: 2,
	},
	"and": {
		Opcode:     and,
		NumOperand: 2,
	},
	"not": {
		Opcode:     not,
		NumOperand: 1,
	},
	"xor": {
		Opcode:     xor,
		NumOperand: 2,
	},
	"or": {
		Opcode:     or,
		NumOperand: 2,
	},
	"shr": {
		Opcode:     shr,
		NumOperand: 2,
	},
	"shl": {
		Opcode:     shl,
		NumOperand: 2,
	},
	"jmp": {
		Opcode:     jmp,
		NumOperand: 1,
	},
	"cmp": {
		Opcode:     cmp,
		NumOperand: 2,
	},
	"je": {
		Opcode:     je,
		NumOperand: 1,
	},
	"jne": {
		Opcode:     jne,
		NumOperand: 1,
	},
	"jz": {
		Opcode:     jz,
		NumOperand: 1,
	},
	"jnz": {
		Opcode:     jnz,
		NumOperand: 1,
	},
	"jge": {
		Opcode:     jge,
		NumOperand: 1,
	},
	"jg": {
		Opcode:     jg,
		NumOperand: 1,
	},
	"jle": {
		Opcode:     jle,
		NumOperand: 1,
	},
	"jl": {
		Opcode:     jl,
		NumOperand: 1,
	},
	"in": {
		Opcode:     in,
		NumOperand: 1,
	},
	"out": {
		Opcode:     out,
		NumOperand: 1,
	},
	"outd": {
		Opcode:     outd,
		NumOperand: 1,
	},
	"end": {
		Opcode:     end,
		NumOperand: 0,
	},
	"call": {
		Opcode:     call,
		NumOperand: 1,
	},
	"ret": {
		Opcode:     ret,
		NumOperand: 0,
	},
	"dn": {
		Opcode:     dn,
		NumOperand: 1,
	},
}

// list of instructions that dont use an operand
var NoOperand = []Opcode{end, ret}

// list of instructions that take 1 operand
var OneOperand = []Opcode{not, jmp, je, jne, jz, jnz, jge, jg, jle, jl, in, out, outd, call, dn}

// list of instructions that take 2 operands
var TwoOperand = []Opcode{mov, add, sub, mul, div, mod, and, xor, or, shr, shl, cmp}

type RegisterType uint8

const (
	r0 RegisterType = iota + 1
	r1
	r2
	r3
	r4
	r5
	r6
	r7

	idx

	pc
)

var RegisterMap = map[string]RegisterType{
	"%r0": r0,
	"%r1": r1,
	"%r2": r2,
	"%r3": r3,
	"%r4": r4,
	"%r5": r5,
	"%r6": r6,
	"%r7": r7,

	"%idx": idx,

	"%pc": pc,
}
