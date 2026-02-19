package main

import (
	"flag"
	"nyantime/decoder"
	"nyantime/runtime"
)

func main() {
	var programPath string
	flag.StringVar(&programPath, "program", "add-3-nums.nyobj", "the assembled file to run")
	flag.Parse()

	contents, err := OpenFile(programPath)

	if err != nil {
		panic(err)
	}

	decoded := decoder.DecodeInstructions(contents)

	err = runtime.Runtime(decoded)

	if err != nil {
		panic(err)
	}
}
