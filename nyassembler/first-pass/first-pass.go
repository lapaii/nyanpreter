package firstpass

import (
	"nyassembler/util"
)

func FirstPass(input []string) (SymbolTable, error) {
	symbolTable := SymbolTable{}

	for idx, line := range input {
		labelMatches := util.LabelRegex.FindString(line)

		if labelMatches == "" {
			continue
		}

		labelMatches = labelMatches[:len(labelMatches)-1] // FindString still returns the trailing ":", this removes it

		// adding 1 so that i can tell when a label
		// is defined on the first line vs not at all
		symbolTable[labelMatches] = idx + 1
	}

	return symbolTable, nil
}
