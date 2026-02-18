package main

import "os"

func WriteFile(outputPath string, data []byte) error {
	err := os.WriteFile(outputPath, data, 0644)

	return err
}
