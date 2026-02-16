package main

import (
	"nyassembler/util"
	"os"
	"strings"
)

func OpenFile(filePath string) ([]string, error) {
	contentBytes, err := os.ReadFile(filePath)

	if err != nil {
		return []string{}, err
	}

	contentString := string(contentBytes)
	lineSplit := strings.Split(contentString, "\n")

	// removing comments and blank likes here instead of doing it twice in every pass
	var output []string

	for _, line := range lineSplit {
		noComments := util.CommentsRegex.ReplaceAllString(line, "")

		if noComments == "" {
			continue
		}

		output = append(output, noComments)
	}

	return output, nil
}
