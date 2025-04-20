package interpreter

import (
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"bufio"
	"os"

	"github.com/chzyer/readline"
)

func TerminalInterpreter() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	lex := lexer.NewLexer()
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		lex.Tokenize(line)
		lex.Debug()
	}
}

func RunFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := bufio.NewScanner(file)
	lex := lexer.NewLexer()
	for buffer.Scan(){
		line := buffer.Text()
		lex.Tokenize(line)
		lex.Debug()
	}
}
