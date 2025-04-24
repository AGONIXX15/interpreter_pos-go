package parser

import (
	"AGONIXX15/interpreter_pos-go.git/ast"
	"AGONIXX15/interpreter_pos-go.git/lexer"
)

type binding_power int

const (
	default_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *Parser) ast.Stmt

type nud_handler func(p *Parser) ast.Expr

type led_handler func(p *Parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler

type nud_lookup map[lexer.TokenKind]nud_handler

type led_lookup map[lexer.TokenKind]led_handler

type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func nud(kind lexer.TokenKind, bp binding_power, nud_fn nud_handler) {
	bp_lu[kind] = bp
	nud_lu[kind] = nud_fn
}

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookUps() {
	nud(lexer.INTEGER, primary, parse_primary_expr)
	nud(lexer.FLOAT, primary, parse_primary_expr)
	nud(lexer.DOUBLE, primary, parse_primary_expr)
	nud(lexer.IDENTIFIER, primary, parse_primary_expr)
	nud(lexer.STRING, primary, parse_primary_expr)
	// additive and multiplicative
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	nud(lexer.DASH, additive, parse_unary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)

	// logical and relational
	led(lexer.EQUAL, relational, parse_binary_expr)
	led(lexer.NOT_EQUAL, relational, parse_binary_expr)
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)
	nud(lexer.NOT, logical, parse_unary_expr)

	nud(lexer.LPAREN, primary, parse_primary_expr)

	stmt(lexer.IDENTIFIER, parse_assignment_stmt)
	stmt(lexer.IF, parse_if_stmt)
	stmt(lexer.FUNCTION, parse_function_stmt)
	stmt(lexer.RETURN, parse_return_stmt)

}
