package main

import (
	"fmt"
	"os"
)

func OpenFile(filePath string) ([]byte, error) {
	contentBytes, err := os.ReadFile(filePath)

	if err != nil {
		return []byte{}, err
	}

	if string(contentBytes[0:5]) != "nya:3" {
		return []byte{}, fmt.Errorf("the input isn't a nyobj file!")
	}

	return contentBytes[5:], nil
}
