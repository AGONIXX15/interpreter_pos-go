package ast

import "fmt"

type Stmt interface {
	stmt()
}

type Expr interface {
	expr()
}

type Type interface {
	_type()
}

func ExpectType[T any](v any) T {
	t, ok := v.(T)
	if !ok {
		panic(fmt.Sprintf("expected type %T, got %T", *new(T), v))
	}
	return t
}

func ExpectExpr[T Expr](exp Expr) T {
	return ExpectType[T](exp)
}

func ExpectStmt[T Stmt](stmt Stmt) T {
	return ExpectType[T](stmt)
}
