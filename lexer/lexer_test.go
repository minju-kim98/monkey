package lexer

import (
	"testing"

	"github.com/minju-kim98/monkey/token"
)

func TestMyLexer(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, outputType := range tests {
		tok := l.NextToken()

		if tok.Type != outputType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected = %q, got = %q", i, outputType.expectedType, tok.Type)
		}

		if tok.Literal != outputType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected = %q, got = %q", i, outputType.expectedLiteral, tok.Literal)
		}
	}
}
