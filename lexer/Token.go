package lexer

import (
	"fmt"
	"slices"
)

type TokenKind int

const (
	EQUAL TokenKind = iota
	NOT_EQUAL
	LESS_EQUAL
	GREATER_EQUAL
	LESS
	GREATER
	AND
	OR
	NOT
	DOUBLE_DOT
	INTEGER
	DOUBLE
	FLOAT
	BOOLEAN
	STRING
	COMMA
	NULL
	IDENTIFIER
	PLUS_ASSIGN
	DASH_ASSIGN
	STAR_ASSIGN
	SLASH_ASSIGN
	ASSIGN
	DOUBLE_STAR
	PLUS
	DASH
	STAR
	SLASH
	LPAREN
	RPAREN
	SEMICOLON
	NEWLINE
	WHITESPACE
)

type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{Kind: kind, Value: value}
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	return slices.Contains(expectedTokens, token.Kind)
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, INTEGER, DOUBLE, FLOAT, STRING, BOOLEAN) {
		fmt.Printf("%s (%s)\n", KindToString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", KindToString(token.Kind))
	}
}

func KindToString(kind TokenKind) string {
	switch kind {
	case EQUAL:
		return "EQUAL"
	case NOT_EQUAL:
		return "NOT_EQUAL"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case NOT:
		return "NOT"
	case DOUBLE_DOT:
		return "DOUBLE_DOT"
	case INTEGER:
		return "INTEGER"
	case DOUBLE:
		return "DOUBLE"
	case FLOAT:
		return "FLOAT"
	case BOOLEAN:
		return "BOOLEAN"
	case STRING:
		return "STRING"
	case COMMA:
		return "COMMA"
	case NULL:
		return "NULL"
	case IDENTIFIER:
		return "IDENTIFIER"
	case PLUS_ASSIGN:
		return "PLUS_ASSIGN"
	case DASH_ASSIGN:
		return "DASH_ASSIGN"
	case STAR_ASSIGN:
		return "STAR_ASSIGN"
	case SLASH_ASSIGN:
		return "SLASH_ASSIGN"
	case ASSIGN:
		return "ASSIGN"
	case DOUBLE_STAR:
		return "DOUBLE_STAR"
	case PLUS:
		return "PLUS"
	case DASH:
		return "DASH"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case SEMICOLON:
		return "SEMICOLON"
	case NEWLINE:
		return "NEWLINE"
	case WHITESPACE:
		return "WHITESPACE"
	default:
		return "UNKNOWN"
	}
}
