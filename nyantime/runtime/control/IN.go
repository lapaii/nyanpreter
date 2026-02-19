package control

import (
	"nyantime/registers"
	"os"
	"shared"

	"golang.org/x/term"
)

func IN(r *registers.Registers, operator shared.Operator, program *[]shared.Instruction) error {
	// switch stdin into 'raw' mode, this hides input
	// https://stackoverflow.com/a/70627571
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	char := make([]byte, 1)

	_, err = os.Stdin.Read(char)
	if err != nil {
		return err
	}

	r.SetAccumulator(int(char[0])) // set accumulator to ascii value of the char entered

	return nil
}
