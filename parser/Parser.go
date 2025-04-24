package parser

import (
	"AGONIXX15/interpreter_pos-go.git/ast"
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"fmt"
)

type Parser struct {
	pos int
	lex *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	createTokenLookUps()
	createTypeTokenLookups()
	return &Parser{
		pos: 0,
		lex: lexer,
	}
}

func (p *Parser) Parse() ast.BlockStmt {
	body := make([]ast.Stmt, 0)
	for {
		token := p.lex.Current()
		if token.Kind == lexer.EOF {
			break
		}
		body = append(body, parse_stmt(p))
	}
	return ast.BlockStmt{
		Body: body,
	}
}

func Debug(block ast.BlockStmt) {
	for line, value := range block.Body {
		fmt.Printf("Stmt %d: %s\n", line+1, type_stmt(value))
	}
}

func type_stmt(stmt ast.Stmt) string {
	switch s := stmt.(type) {
	case ast.IfStmt:
		return fmt.Sprintf("IF(branches=%v,Else=%v)", s.Branches, s.Else)
	case ast.AssignmentStmt:
		return fmt.Sprintf("ASSIGNMENT(Name=%s, Value=%v, type=%v, inferred=%v)", s.Name, s.Value, s.ExplicitType, s.Inferred)

	case ast.FunctionStmt:
		return fmt.Sprintf("FUNCTION(Name=%s, Parameters=%v,TypeReturn=%v, Body=%v)", s.Name, s.Parameters, s.TypeReturn, s.Body)
	default:
		return "UNKOWN"
	}
}
