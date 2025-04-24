package ast

import "AGONIXX15/interpreter_pos-go.git/lexer"

type Integer struct {
	Value int
}

func (number Integer) expr() {}

type Float struct {
	Value float32
}

func (number Float) expr() {}

type Double struct {
	Value float64
}

func (number Double) expr() {}

type String struct {
	Value string
}

func (str String) expr() {}

type Symbol struct {
	Value string
}

func (str Symbol) expr() {}


// composite Expr

type BinaryExpr struct {
	Left Expr
	OperatorToken lexer.Token
	Right Expr
}

func (bE BinaryExpr) expr() {}

type UnaryExpr struct {
	OperatorToken lexer.Token
	Right Expr
}

func (uE UnaryExpr) expr() {}


