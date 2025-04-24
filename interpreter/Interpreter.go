package interpreter

import (
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"AGONIXX15/interpreter_pos-go.git/parser"
	"bufio"
	"os"
	// "fmt"

	// "github.com/chzyer/readline"
)

// func TerminalInterpreter() {
// 	buffer := bufio.NewScanner(os.Stdin)
// 	lex := lexer.NewLexer(buffer)
// 	p := parser.NewParser(lex)
// 	if buffer.Scan() {
// 		line := buffer.Text()
// 		fmt.Println(line)
// 		body := p.Parse()
// 		fmt.Printf("%v\n",body)
// 		parser.Debug(body)
// 	}
// }



func RunFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := bufio.NewScanner(file)
	lex := lexer.NewLexer(buffer)
	p := parser.NewParser(lex)
	block := p.Parse()
	parser.Debug(block)
}
