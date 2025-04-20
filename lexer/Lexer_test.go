package lexer_test

import (
	"AGONIXX15/interpreter_pos-go.git/lexer"
	"reflect"
	"testing"
)

func TestLexer_Tokenize(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line string
		want []lexer.Token
	}{
		// TODO: Add test cases.
		{name: "sumas",
			line: "10 + 3",
			want: []lexer.Token{{lexer.INTEGER, "10"}, {lexer.PLUS, "+"},
			{lexer.INTEGER, "3"}},
		},
		{
			name: "concatenacion",
			line: "\"hola\" +\"mundo\"",
			want: []lexer.Token{{lexer.STRING, "hola"}, {lexer.PLUS, "+"}, {lexer.STRING, "mundo"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := lexer.NewLexer()
			got := lexer.Tokenize(tt.line)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
