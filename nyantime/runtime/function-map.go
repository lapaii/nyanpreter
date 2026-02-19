package runtime

import (
	"nyantime/registers"
	"nyantime/runtime/bitwise"
	"nyantime/runtime/control"
	"nyantime/runtime/math"
	"nyantime/runtime/memory"
	"nyantime/util"
)

var FunctionMap = map[util.Operand]func(*registers.Registers, util.Operator, *[]util.Instruction) error{
	util.LDM: memory.LDM,
	util.LDD: memory.LDD,
	util.LDI: memory.LDI,
	util.LDX: memory.LDX,
	util.LDR: memory.LDR,
	util.MOV: memory.MOV,
	util.STO: memory.STO,

	util.ADD: math.ADD,
	util.SUB: math.SUB,
	util.INC: math.INC,
	util.DEC: math.DEC,

	util.JMP: control.JMP,
	util.CMP: control.CMP,
	util.CMI: control.CMI,
	util.JE:  control.JE,
	util.JNE: control.JNE,
	util.IN:  control.IN,
	util.OUT: control.OUT,

	util.AND: bitwise.AND,
	util.XOR: bitwise.XOR,
	util.OR:  bitwise.OR,
	util.LSL: bitwise.LSL,
	util.LSR: bitwise.LSR,
}
