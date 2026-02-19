package util

import (
	"strconv"
)

func ParseOperator(operator Operator) int {
	numberBase := 10

	switch operator[0] {
	case '#':
		numberBase = 10
	case 'B':
		numberBase = 2
	case '&':
		numberBase = 16
	}

	conv, err := strconv.ParseInt(string(operator)[1:], numberBase, 0)

	if err != nil {
		return 0
	}

	return int(conv)
}
