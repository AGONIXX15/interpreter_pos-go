package parser

import (
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"fmt"
)

func (p *Parser) getToken() lexer.Token {
	return p.lex.Current()
}

func (p *Parser) getCurrentKind() lexer.TokenKind {
	return p.lex.Current().Kind
}

func (p *Parser) nextToken() lexer.Token {
	return p.lex.Next()
}

func (p *Parser) expect(kind lexer.TokenKind, err any) {
	current_token := p.getToken()
	if current_token.Kind == kind {
		p.nextToken()
		return
	}
	if err == nil {
		panic(fmt.Sprintf("expected at line %d column %d %v got %v", p.lex.Line, p.lex.Column, lexer.KindToString(kind), lexer.KindToString(current_token.Kind)))
	}
	panic(err)
}
