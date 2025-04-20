package parser

import (
	"AGONIXX15/interpreter_pos-go.git/lexer"
)


type Parser struct {
	pos int
	tokens []lexer.Token
	lex *lexer.Lexer
}
