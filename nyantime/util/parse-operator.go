package util

import (
	"shared"
	"strconv"
)

func ParseOperator(operator shared.Operator) (int, error) {
	numberBase := -1

	switch operator[0] {
	case '#':
		numberBase = 10
	case 'B':
		numberBase = 2
	case '&':
		numberBase = 16
	}

	// operator is a label, return just the line number it means, i will load that in the instructions
	if numberBase == -1 {
		conv, err := strconv.Atoi(string(operator))

		if err != nil {
			return 0, err
		}

		return conv, nil
	}

	conv, err := strconv.ParseInt(string(operator)[1:], numberBase, 0)

	if err != nil {
		return 0, err
	}

	return int(conv), nil
}
