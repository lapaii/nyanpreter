package parser

import (
	"fmt"
	"os"
	"strings"
)

func LoadFile(path string) ([]string, error) {
	contentBytes, err := os.ReadFile(path)

	if err != nil {
		return []string{}, fmt.Errorf("Failed to read program file, reason: %s", err.Error())
	}

	content := string(contentBytes)

	lines := strings.Split(content, "\n")

	return lines, nil
}
