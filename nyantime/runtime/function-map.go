package runtime

import (
	"nyantime/registers"
	"nyantime/runtime/control"
	"nyantime/runtime/math"
	"nyantime/runtime/memory"
	"nyantime/util"
)

var FunctionMap = map[util.Operand]func(*registers.Registers, util.Operator, *[]util.Instruction) error{
	util.LDM: memory.LDM,
	util.LDD: memory.LDD,
	util.STO: memory.STO,

	util.ADD: math.ADD,

	util.OUT: control.OUT,
}
