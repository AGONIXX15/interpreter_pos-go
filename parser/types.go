package parser

import (
	"AGONIXX15/interpreter_pos-go.git/ast"
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"fmt"
)

type type_led_handler func(p *Parser, left ast.Type, bp binding_power) ast.Type
type type_nud_handler func(p *Parser) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = type_bp_lookup{}
var type_led_lu = type_led_lookup{}
var type_nud_lu = type_nud_lookup{}

func type_nud(kind lexer.TokenKind, bp binding_power, nud_fn type_nud_handler) {
	type_bp_lu[kind] = bp
	type_nud_lu[kind] = nud_fn

}
func type_led(kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func createTypeTokenLookups() {
	type_nud(lexer.IDENTIFIER, primary, func(p *Parser) ast.Type {
		return ast.SymbolType{Value: p.nextToken().Value}
	})

	type_nud(lexer.LBRACKET, member, func(p *Parser) ast.Type {
		p.expect(lexer.LBRACKET, nil)
		p.expect(lexer.RBRACKET, nil)
		value_type := parse_type(p, default_bp)
		return ast.ListType{
			Value: value_type,
		}
	})

}

func parse_type(p *Parser, bp binding_power) ast.Type {
	token := p.getToken()
	nud_fn, exist := type_nud_lu[token.Kind]
	if !exist {
		panic(fmt.Sprintf("type: NUD handler expected for token %v", lexer.KindToString(token.Kind)))
	}
	left := nud_fn(p)
	// fmt.Printf("type_bp_lu[] = %v >(%v) bp", lexer.KindToString(p.getCurrentKind()), bp)
	for p.getCurrentKind() != lexer.EOF && type_bp_lu[p.getCurrentKind()] > bp {
		token = p.getToken()
		fmt.Printf("%v\n", type_led_lu)
		led_fn, exist := type_led_lu[token.Kind]
		fmt.Println(exist)
		if !exist {
			panic(fmt.Sprintf("type: led handler expected for token %v", lexer.KindToString(token.Kind)))
		}
		left = led_fn(p, left, bp)
	}
	return left
}
