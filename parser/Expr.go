package parser

import (
	"AGONIXX15/interpreter_pos-go.git/ast"
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"fmt"
	"strconv"
)

func parse_expr(p *Parser, bp binding_power) ast.Expr {
	kind := p.getCurrentKind()
	nud_fn, exist := nud_lu[kind]
	if !exist {
		panic(fmt.Sprintf("expected nud token %s", lexer.KindToString(kind)))
	}
	left := nud_fn(p)
	for p.getToken().Kind != lexer.EOF && bp < bp_lu[p.getCurrentKind()] {
		led_fn, exist := led_lu[p.getCurrentKind()]
		if !exist {
			panic(fmt.Sprintf("expected led token %s", lexer.KindToString(p.getCurrentKind())))
		}
		left = led_fn(p, left, bp)
	}
	return left
}

func parse_primary_expr(p *Parser) ast.Expr {
	token := p.nextToken()
	switch token.Kind {
	case lexer.INTEGER:
		integer, _ := strconv.Atoi(token.Value)
		return ast.Integer{Value: integer}
	case lexer.FLOAT:
		fl64, _ := strconv.ParseFloat(token.Value, 32)
		fl := float32(fl64)
		return ast.Float{Value: fl}
	case lexer.DOUBLE:
		fl, _ := strconv.ParseFloat(token.Value, 64)
		return ast.Double{Value: fl}
	case lexer.STRING:
		return ast.String{Value: token.Value}
	case lexer.IDENTIFIER:
		return ast.Symbol{Value: token.Value}
	case lexer.LPAREN:
		expr := parse_expr(p, default_bp)
		p.expect(lexer.RPAREN, nil)
		return expr
	default:
		panic(fmt.Sprintf("it is not a primary expression %v", lexer.KindToString(token.Kind)))
	}
}

func parse_unary_expr(p *Parser) ast.Expr {
	operator := p.nextToken() // consume op
	right := parse_expr(p, bp_lu[operator.Kind])
	return ast.UnaryExpr{OperatorToken: operator, Right: right}

}

func parse_binary_expr(p *Parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.nextToken() // consume op
	right := parse_expr(p, bp_lu[operator.Kind])
	return ast.BinaryExpr{Left: left, OperatorToken: operator, Right: right}
}
