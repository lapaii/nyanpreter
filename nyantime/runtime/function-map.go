package runtime

import (
	"nyantime/registers"
	"nyantime/runtime/bitwise"
	"nyantime/runtime/control"
	"nyantime/runtime/math"
	"nyantime/runtime/memory"
	"shared"
)

var FunctionMap = map[shared.Operand]func(*registers.Registers, shared.Operator, *[]shared.Instruction) error{
	shared.LDM: memory.LDM,
	shared.LDD: memory.LDD,
	shared.LDI: memory.LDI,
	shared.LDX: memory.LDX,
	shared.LDR: memory.LDR,
	shared.MOV: memory.MOV,
	shared.STO: memory.STO,

	shared.ADD: math.ADD,
	shared.SUB: math.SUB,
	shared.INC: math.INC,
	shared.DEC: math.DEC,

	shared.JMP: control.JMP,
	shared.CMP: control.CMP,
	shared.CMX: control.CMX,
	shared.CMI: control.CMI,
	shared.JE:  control.JE,
	shared.JNE: control.JNE,
	shared.JGE: control.JGE,
	shared.JLE: control.JLE,
	shared.IN:  control.IN,
	shared.OUT: control.OUT,

	shared.AND: bitwise.AND,
	shared.XOR: bitwise.XOR,
	shared.OR:  bitwise.OR,
	shared.LSL: bitwise.LSL,
	shared.LSR: bitwise.LSR,
}
