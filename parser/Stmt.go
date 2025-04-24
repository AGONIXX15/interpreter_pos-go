package parser

import (
	"AGONIXX15/interpreter_pos-go.git/ast"
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"fmt"
)

func parse_stmt(p *Parser) ast.Stmt {
	stmt_fn, exist := stmt_lu[p.getCurrentKind()] // assignment
	if !exist {
		panic(fmt.Sprintf("not a statement %v", lexer.KindToString(p.getCurrentKind())))
	}
	return stmt_fn(p)
}

func parse_expression_stmt(p *Parser) ast.ExpressionStmt {
	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON, nil)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_assignment_stmt(p *Parser) ast.Stmt {
	name := p.nextToken().Value
	if p.getCurrentKind() == lexer.COLON {
		p.expect(lexer.COLON, nil)
		_type := parse_type(p, default_bp)
		if p.getCurrentKind() == lexer.ASSIGN {
			p.expect(lexer.ASSIGN, nil)
			value := parse_expr(p, default_bp)
			p.expect(lexer.SEMICOLON, nil)
			return ast.AssignmentStmt{Name: name, ExplicitType: _type, Value: value, Inferred: false}
		} else {
			p.expect(lexer.ASSIGN, nil)
		}
	}

	if p.getCurrentKind() == lexer.COLON_ASSIGN {
		p.expect(lexer.COLON_ASSIGN, nil)
		value := parse_expr(p, default_bp)
		p.expect(lexer.SEMICOLON, nil)
		return ast.AssignmentStmt{Name: name, ExplicitType: nil, Value: value, Inferred: true}
	}
	panic(fmt.Sprintf("expected = or := not founded"))
}

func parse_block_stmt(p *Parser) ast.Stmt {
	p.expect(lexer.LBRACE, nil)
	stmts := make([]ast.Stmt, 0)
	for p.getCurrentKind() != lexer.RBRACE {
		stmts = append(stmts, parse_stmt(p))
	}
	p.expect(lexer.RBRACE, nil)
	return ast.BlockStmt{Body: stmts}
}

func parse_if_stmt(p *Parser) ast.Stmt {
	p.expect(lexer.IF, nil)
	condition := parse_expr(p, default_bp)
	body := parse_block_stmt(p)
	branches := []ast.IfCondition{
		{Condition: condition, Body: ast.ExpectStmt[ast.BlockStmt](body)},
	}
	for p.getCurrentKind() == lexer.ELIF {
		p.expect(lexer.ELIF, nil)
		cond := parse_expr(p, default_bp)
		blk := parse_block_stmt(p)
		branches = append(branches, ast.IfCondition{
			Condition: cond, Body: ast.ExpectStmt[ast.BlockStmt](blk),
		})
	}
	var elseBody ast.BlockStmt
	if p.getCurrentKind() == lexer.ELSE {
		p.expect(lexer.ELSE, nil)
		blk := parse_block_stmt(p)
		elseBody = ast.ExpectStmt[ast.BlockStmt](blk)
	}

	return ast.IfStmt{
		Branches: branches,
		Else:     elseBody,
	}
}

func parse_parameter(p *Parser) ast.Parameter {
	name := p.nextToken().Value
	p.expect(lexer.COLON, nil)
	_type := parse_type(p, default_bp)
	return ast.Parameter{Name: name, TypeParameter: _type}
}

func parse_parameters(p *Parser) []ast.Parameter {
	p.expect(lexer.LPAREN, nil)
	parameters := make([]ast.Parameter, 0)
	for p.getCurrentKind() != lexer.RPAREN {
		if p.getCurrentKind() != lexer.IDENTIFIER {
			panic(fmt.Sprintf("expected IDENTIFIER got %s", lexer.KindToString(p.getCurrentKind())))
		}
		// fmt.Printf("algo: %v", p.getToken())
		parameters = append(parameters, parse_parameter(p))
		if p.getCurrentKind() == lexer.COMMA {
			p.expect(lexer.COMMA, nil)
		}
	}
	p.expect(lexer.RPAREN, nil)
	return parameters
}

func parse_function_stmt(p *Parser) ast.Stmt {
	p.expect(lexer.FUNCTION, nil)
	if p.getCurrentKind() != lexer.IDENTIFIER {
		panic(fmt.Sprintf("expected IDENTIFIER got %s", lexer.KindToString(p.getCurrentKind())))
	}
	name := p.nextToken().Value
	parameters := parse_parameters(p)
	_type := parse_type(p, default_bp)
	body := parse_block_stmt(p)
	return ast.FunctionStmt{
		Name:       name,
		Parameters: parameters,
		TypeReturn: _type,
		Body:       ast.ExpectStmt[ast.BlockStmt](body),
	}
}

func parse_return_stmt(p *Parser) ast.Stmt {
	p.expect(lexer.RETURN, nil)
	value := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON, nil)
	return ast.ReturnStmt{Value: value}
}
