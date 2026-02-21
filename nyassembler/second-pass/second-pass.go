package secondpass

import (
	"fmt"
	firstpass "nyassembler/first-pass"
	"nyassembler/second-pass/operand"
	"nyassembler/util"
	"shared"
	"strings"
)

func SecondPass(contents []string, symbolTable firstpass.SymbolTable) ([]shared.Instruction, error) {
	var outputProgram = []shared.Instruction{}

	for idx, line := range contents {
		line := strings.ReplaceAll(line, ",", "")

		lineParts := strings.Fields(line)

		// if line is a label definition, remove label definition as we dont need it here
		if util.LabelRegex.MatchString(lineParts[0]) {
			lineParts = lineParts[1:]
		}

		OpcodeInfo := shared.InstructionSet[lineParts[0]]

		// invalid opcode!
		if OpcodeInfo.Opcode == 0 {
			return []shared.Instruction{}, fmt.Errorf("[line %d] invalid opcode found! %s", idx+1, lineParts[0])
		}

		// check if correct amount of operands were received
		// if not, return error
		if len(lineParts) != int(OpcodeInfo.NumOperand)+1 {
			return []shared.Instruction{}, fmt.Errorf("[line %d] %d operand(s) expected, received %d", idx+1, OpcodeInfo.NumOperand, len(lineParts)-1)
		}

		// init operand stuff
		var src shared.Operand
		var dest shared.Operand
		var srcType shared.OperandType
		var destType shared.OperandType

		// parse first operand
		var err error
		if OpcodeInfo.NumOperand >= 1 {
			dest, destType, err = operand.ParseOperand(idx, lineParts[1], symbolTable)

			if err != nil {
				return []shared.Instruction{}, err
			}
		}

		// parse second operand
		if OpcodeInfo.NumOperand == 2 {
			src, srcType, err = operand.ParseOperand(idx, lineParts[2], symbolTable)

			if err != nil {
				return []shared.Instruction{}, err
			}
		}

		// add this instruction
		outputProgram = append(outputProgram, shared.Instruction{
			SourceType: srcType,
			DestType:   destType,
			Opcode:     OpcodeInfo.Opcode,
			Source:     src,
			Dest:       dest,
		})
	}

	return outputProgram, nil
}
