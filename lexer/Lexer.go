package lexer

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type RegexHandler func(lexer *Lexer, regex *regexp.Regexp)

type RegexPattern struct {
	Regex   *regexp.Regexp
	Handler RegexHandler
}

type Lexer struct {
	File        *bufio.Scanner
	pos         int
	Line        int
	Column      int
	currentLine string
	Tokens      []Token
	Patterns    []RegexPattern
}

func (lexer *Lexer) advance(n int) {
	lexer.pos += n
	lexer.Column += n
}

func (lexer *Lexer) push(token Token) {
	lexer.Tokens = append(lexer.Tokens, token)
}

func (lexer *Lexer) remainder() string {
	return lexer.currentLine[lexer.pos:]
}

func (lexer *Lexer) Next() Token {
	if len(lexer.Tokens) == 0 {
		lexer.Tokenize()
	}
	token := lexer.Tokens[0]
	lexer.Tokens = lexer.Tokens[1:len(lexer.Tokens)]
	return token
}

func (lexer *Lexer) Current() Token {
	if len(lexer.Tokens) == 0 {
		lexer.Tokenize()
	}
	token := lexer.Tokens[0]
	return token
}

func (lexer *Lexer) reset() {
	lexer.pos = 0
	lexer.Tokens = make([]Token, 0)
	lexer.Line += 1
	lexer.Column = 1
}

func NewLexer(file *bufio.Scanner) *Lexer {
	return &Lexer{
		File:        file,
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
			{regexp.MustCompile("\\+="), defaultHandler(PLUS_ASSIGN, "+=")},
			{regexp.MustCompile("-="), defaultHandler(DASH_ASSIGN, "-=")},
			{regexp.MustCompile("\\*="), defaultHandler(STAR_ASSIGN, "*=")},
			{regexp.MustCompile("/="), defaultHandler(SLASH_ASSIGN, "/=")},
			{regexp.MustCompile("\\*\\*"), defaultHandler(DOUBLE_STAR, "**")},
			{regexp.MustCompile("&&"), defaultHandler(AND, "&&")},
			{regexp.MustCompile("\\|\\|"), defaultHandler(OR, "||")},
			{regexp.MustCompile("func"), defaultHandler(FUNCTION, "func")},
			{regexp.MustCompile("if"), defaultHandler(IF, "if")},
			{regexp.MustCompile("elif"), defaultHandler(ELIF, "elif")},
			{regexp.MustCompile("else"), defaultHandler(ELSE, "else")},
			{regexp.MustCompile("end"), defaultHandler(END, "end")},
			{regexp.MustCompile("return"), defaultHandler(RETURN, "return")},
			{regexp.MustCompile("true|false"), valueHandler(BOOLEAN)},
			{regexp.MustCompile("null"), valueHandler(NULL)},
			{regexp.MustCompile("[0-9]+\\.[0-9]*f"), valueHandler(FLOAT)},
			{regexp.MustCompile("[0-9]+\\.[0-9]+"), valueHandler(DOUBLE)},
			{regexp.MustCompile("[0-9]+"), valueHandler(INTEGER)},
			{regexp.MustCompile(`^"(\\.|[^"\\])*"`), stringHandler()},
			{regexp.MustCompile("[a-zA-Z_][a-zA-Z0-9_]*"), valueHandler(IDENTIFIER)},
			{regexp.MustCompile("!"), defaultHandler(NOT, "!")},
			{regexp.MustCompile("<"), defaultHandler(LESS, "<")},
			{regexp.MustCompile(">"), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(":="), defaultHandler(COLON_ASSIGN, ":=")},
			{regexp.MustCompile("="), defaultHandler(ASSIGN, "=")},
			{regexp.MustCompile("\\+"), defaultHandler(PLUS, "+")},
			{regexp.MustCompile("-"), defaultHandler(DASH, "-")},
			{regexp.MustCompile("\\*"), defaultHandler(STAR, "*")},
			{regexp.MustCompile("/"), defaultHandler(SLASH, "/")},
			{regexp.MustCompile("\\.\\."), defaultHandler(DOUBLE_DOT, "..")},
			{regexp.MustCompile(","), defaultHandler(COMMA, ",")},
			{regexp.MustCompile("\\("), defaultHandler(LPAREN, "(")},
			{regexp.MustCompile("\\)"), defaultHandler(RPAREN, ")")},
			{regexp.MustCompile("\\["), defaultHandler(LBRACKET, "[")},
			{regexp.MustCompile("\\]"), defaultHandler(RBRACKET, "]")},
			{regexp.MustCompile("{"), defaultHandler(LBRACE, "{")},
			{regexp.MustCompile("}"), defaultHandler(RBRACE, "}")},
			{regexp.MustCompile(":"), defaultHandler(COLON, ":")},
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
func stringHandler() RegexHandler {
	return func(lexer *Lexer, regex *regexp.Regexp) {
		value := regex.FindString(lexer.remainder())
		lexer.advance(len(value))
		value = value[1 : len(value)-1]
		lexer.push(NewToken(STRING, value))
	}
}

func trashHandler() RegexHandler {
	return func(lexer *Lexer, regex *regexp.Regexp) {
		value := regex.FindString(lexer.remainder())
		lexer.advance(len(value))
	}
}

func (lexer *Lexer) Tokenize() {
	if lexer.File.Scan() {
		line := lexer.File.Text()
		if strings.TrimRight(line, " ") == "" {
			lexer.Tokenize()
			return
		}
		lexer.currentLine = line
	} else {
		lexer.Tokens = append(lexer.Tokens, Token{EOF, ""})
		return
	}
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
}

func (lexer *Lexer) Debug() {
	fmt.Printf("line %d: %v\n", lexer.Line, lexer.currentLine)
	for _, token := range lexer.Tokens {
		token.Debug()
	}
}
