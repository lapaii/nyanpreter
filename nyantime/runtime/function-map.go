package runtime

import (
	"nyantime/registers"
	"nyantime/runtime/memory"
	"nyantime/util"
)

var FunctionMap = map[util.Operand]func(*registers.Registers, util.Operator, []util.Instruction) error{
	util.LDM: memory.LDM,
	util.LDD: memory.LDD,
}
