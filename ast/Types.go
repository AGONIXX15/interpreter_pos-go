package ast

type SymbolType struct {
	Value string
}

func (s SymbolType) _type() {}

type ListType struct {
	Value Type
}

func (l ListType) _type() {}
