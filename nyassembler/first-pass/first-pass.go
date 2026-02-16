package firstpass

import (
	"nyassembler/util"
)

func FirstPass(input []string) ([]SymbolTable, error) {
	symbolTable := []SymbolTable{}

	for idx, line := range input {
		labelMatches := util.LabelRegex.FindString(line)

		if labelMatches == "" {
			continue
		}

		labelMatches = labelMatches[:len(labelMatches)-1] // FindString still returns the trailing ":", this removes it

		symbolTable = append(symbolTable, SymbolTable{
			LabelName: labelMatches,
			LineNum:   idx,
		})
	}

	return symbolTable, nil
}
