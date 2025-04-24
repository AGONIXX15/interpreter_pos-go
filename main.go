package main

import (
	"AGONIXX15/interpreter_pos-go.git/interpreter"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// interpreter.TerminalInterpreter()
		return
	}

	for _, file := range os.Args[1:] {
		interpreter.RunFile(file)
	}
}
