package lexer

import (
	"fmt"
	"regexp"
)

type RegexHandler func(lexer *Lexer, regex *regexp.Regexp)

type RegexPattern struct {
	Regex   *regexp.Regexp
	Handler RegexHandler
}

type Lexer struct {
	pos         int
	Line        int
	Column      int
	currentLine string
	Tokens      []Token
	Patterns    []RegexPattern
}

func (lexer *Lexer) advance(n int) {
	lexer.pos += n
}

func (lexer *Lexer) push(token Token) {
	lexer.Tokens = append(lexer.Tokens, token)
}

func (lexer *Lexer) remainder() string {
	return lexer.currentLine[lexer.pos:]
}

func (lexer *Lexer) reset() {
	lexer.pos = 0
	lexer.Tokens = make([]Token, 0)
	lexer.Line += 1
	lexer.Column = 1
}
func NewLexer() *Lexer {
	return &Lexer{
		pos:         0,
		Line:        0,
		Column:      1,
		currentLine: "",
		Tokens:      make([]Token, 0),
		Patterns: []RegexPattern{
			{regexp.MustCompile("=="), defaultHandler(EQUAL, "==")},
			{regexp.MustCompile("!="), defaultHandler(NOT_EQUAL, "!=")},
			{regexp.MustCompile("<="), defaultHandler(LESS_EQUAL, "<=")},
			{regexp.MustCompile(">="), defaultHandler(GREATER_EQUAL, ">=")},
			{regexp.MustCompile("<"), defaultHandler(LESS, "<")},
			{regexp.MustCompile(">"), defaultHandler(GREATER, ">")},
			{regexp.MustCompile("&&"), defaultHandler(AND, "&&")},
			{regexp.MustCompile("\\|\\|"), defaultHandler(OR, "||")},
			{regexp.MustCompile("!"), defaultHandler(NOT, "!")},
			{regexp.MustCompile("\\.\\."), defaultHandler(DOUBLE_DOT, "..")},

			{regexp.MustCompile("[0-9]+\\.[0-9]*f"), valueHandler(FLOAT)},
			{regexp.MustCompile("[0-9]+\\.[0-9]+"), valueHandler(DOUBLE)},
			{regexp.MustCompile("[0-9]+"), valueHandler(INTEGER)},
			{regexp.MustCompile("true|false"), valueHandler(BOOLEAN)},
			{regexp.MustCompile(`^"(\\.|[^"\\])*"`), valueHandler(STRING)},
			{regexp.MustCompile(","), defaultHandler(COMMA, ",")},
			{regexp.MustCompile("null"), valueHandler(NULL)},
			{regexp.MustCompile("[a-zA-Z_][a-zA-Z0-9_]*"), valueHandler(IDENTIFIER)},
			{regexp.MustCompile("\\+="), defaultHandler(PLUS_ASSIGN, "+=")},
			{regexp.MustCompile("-="), defaultHandler(DASH_ASSIGN, "-=")},
			{regexp.MustCompile("\\*="), defaultHandler(STAR_ASSIGN, "*=")},
			{regexp.MustCompile("/="), defaultHandler(SLASH_ASSIGN, "/=")},
			{regexp.MustCompile("="), defaultHandler(ASSIGN, "=")},
			{regexp.MustCompile("\\*\\*"), defaultHandler(DOUBLE_STAR, "**")},
			{regexp.MustCompile("\\+"), defaultHandler(PLUS, "+")},
			{regexp.MustCompile("-"), defaultHandler(DASH, "-")},
			{regexp.MustCompile("\\*"), defaultHandler(STAR, "*")},
			{regexp.MustCompile("/"), defaultHandler(SLASH, "/")},
			{regexp.MustCompile("\\("), defaultHandler(LPAREN, "(")},
			{regexp.MustCompile("\\)"), defaultHandler(RPAREN, ")")},
			{regexp.MustCompile(";"), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile("\\n"), trashHandler()},
			{regexp.MustCompile(" "), trashHandler()},
		},
	}
}

func defaultHandler(kind TokenKind, value string) RegexHandler {
	return func(lexer *Lexer, regex *regexp.Regexp) {
		lexer.advance(len(value))
		lexer.push(NewToken(kind, value))
	}
}

func valueHandler(kind TokenKind) RegexHandler {
	return func(lexer *Lexer, regex *regexp.Regexp) {
		value := regex.FindString(lexer.remainder())
		lexer.advance(len(value))
		lexer.push(NewToken(kind, value))
	}
}

func trashHandler() RegexHandler {
	return func(lexer *Lexer, regex *regexp.Regexp) {
		value := regex.FindString(lexer.remainder())
		lexer.advance(len(value))
	}
}

func (lexer *Lexer) Tokenize(line string) []Token {
	lexer.currentLine = line
	lexer.reset()
	n := len(lexer.currentLine)
	for lexer.pos < n {
		matched := false

		for _, pattern := range lexer.Patterns {
			loc := pattern.Regex.FindStringIndex(lexer.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.Handler(lexer, pattern.Regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("lexer::error line: %d, column: %d unrecognized value: %v", lexer.Line, lexer.Column, lexer.remainder()))
		}
	}
	return lexer.Tokens
}

func (lexer *Lexer) Debug() {
	fmt.Printf("line %d: %v\n", lexer.Line, lexer.currentLine)
	for _, token := range lexer.Tokens {
		token.Debug()
	}
}
