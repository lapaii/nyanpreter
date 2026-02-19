package secondpass

import (
	"fmt"
	firstpass "nyassembler/first-pass"
	"regexp"
	"shared"
	"slices"
	"strconv"
	"strings"
)

// function is rather "messy", should get cleaned up some time
func SecondPass(contents []string, symbolTable firstpass.SymbolTable) ([]shared.Instruction, error) {
	var outputProgram []shared.Instruction

	numberOperator := regexp.MustCompile(`(#[0-9]+)|(B[0-1]+)|(\^[0-9A-F]+)`)
	rawNumberOperator := regexp.MustCompile(`([0-9]+)`)

	for idx, line := range contents {
		// stuff this step does
		// check if operand is valid, if not: throw error and stop
		// check if operator is a label, if so: replace with value from symbol table

		// going to keep storing operator as a string, this is so that i can
		// encode numbers as #n/Bn/&n and addresses as just `n`

		// if the operator isn't compatible with the operand, throw an error and stop

		// output

		lineParts := strings.Fields(line)

		// operator will always be 1 more than operand
		// setting to 0 so that i can tell when something goes wrong
		var operandIndex = -1

		switch len(lineParts) {
		// label definition lines
		case 3:
			operandIndex = 1
		// operand operator or just operand lines
		case 2, 1:
			operandIndex = 0
		}

		// something has gone wrong, throw error!!
		if operandIndex == -1 {
			return []shared.Instruction{}, fmt.Errorf("something went wrong on line %d!! %s", idx, line)
		}

		operand, err := ParseOperand(lineParts[operandIndex], idx)

		if err != nil {
			return []shared.Instruction{}, err
		}

		// operand DOESNT get an operator
		if slices.Contains(shared.NoOperator, operand) {
			// no operator in the line
			if operandIndex+1 > len(lineParts)-1 {
				outputProgram = append(outputProgram, shared.Instruction{
					Operand:  operand,
					Operator: "",
				})

				continue
			}

			return []shared.Instruction{}, fmt.Errorf("operator with operand that doesn't support an operator! on line %d %s", idx, line)
		}

		// dealt with operands without operators, now just throw an error if there isn't one
		if operandIndex+1 > len(lineParts)-1 {
			return []shared.Instruction{}, fmt.Errorf("no operator found on line %d! %s", idx, line)
		}

		operator := lineParts[operandIndex+1]

		// operand requires a register for the operator (ACC or IDX)
		if slices.Contains(shared.RegisterOperator, operand) {
			// is it a register?
			if operator != "ACC" && operator != "IDX" {
				return []shared.Instruction{}, fmt.Errorf("operator on line %d isn't a register! %s", idx, line)
			}

			outputProgram = append(outputProgram, shared.Instruction{
				Operand:  operand,
				Operator: shared.Operator(operator),
			})

			continue
		}

		// now we are dealing with only instructions
		// that use numbers / addresses (or labels to represent either of those)

		// instruction can only take defined numbers
		if slices.Contains(shared.NumberOperator, operand) {
			// this operator is a number
			if numberOperator.MatchString(operator) {
				outputProgram = append(outputProgram, shared.Instruction{
					Operand:  operand,
					Operator: shared.Operator(operator),
				})

				continue
			}

			return []shared.Instruction{}, fmt.Errorf("operator on line %d isn't a defined number! %s", idx, line)
		}

		// instruction can only take addresses
		if slices.Contains(shared.AddressOperator, operand) {
			// operator is a defined number, not allowed
			if numberOperator.MatchString(operator) {
				return []shared.Instruction{}, fmt.Errorf("operator on line %d isn't an address! %s", idx, line)
			}

			// operator is a "raw" address, i.e. LDD 5 <- load operator at address 5
			if rawNumberOperator.MatchString(operator) {
				// just add to output
				outputProgram = append(outputProgram, shared.Instruction{
					Operand:  operand,
					Operator: shared.Operator(operator),
				})

				continue
			}

			// operator must be a label now
			symbolTableLine := symbolTable[operator]

			// operator isnt in symbol table
			if symbolTableLine == 0 {
				return []shared.Instruction{}, fmt.Errorf("%s is used as a symbol but isn't defined at all! line %d", operator, idx)
			}

			// label is properly defined and used
			outputProgram = append(outputProgram, shared.Instruction{
				Operand:  operand,
				Operator: shared.Operator(strconv.Itoa(symbolTableLine - 1)),
			})

			continue
		}

		// instruction can take either user defined number or an address
		if numberOperator.MatchString(operator) {
			outputProgram = append(outputProgram, shared.Instruction{
				Operand:  operand,
				Operator: shared.Operator(operator),
			})

			continue
		}

		// not a defined number, at this point it must be an address
		// operator is a "raw" address
		if rawNumberOperator.MatchString(operator) {
			// just add to output
			outputProgram = append(outputProgram, shared.Instruction{
				Operand:  operand,
				Operator: shared.Operator(operator),
			})

			continue
		}

		// must be a label now
		symbolTableLine := symbolTable[operator]

		// operator isnt in symbol table
		if symbolTableLine == 0 {
			return []shared.Instruction{}, fmt.Errorf("%s is used as a symbol but isn't defined at all! line %d", operator, idx)
		}

		// label is properly defined and used
		outputProgram = append(outputProgram, shared.Instruction{
			Operand:  operand,
			Operator: shared.Operator(strconv.Itoa(symbolTableLine - 1)),
		})

		continue
	}

	return outputProgram, nil
}
