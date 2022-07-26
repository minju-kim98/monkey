package lexer

import (
	"testing"

	"github.com/minju-kim98/monkey/token"
)

func TestBasicToken(t *testing.T) {
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

func TestLetToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
	  x + y;
	};
	
	let result = add(five, ten);

	let my_test$123 = 15;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.ID, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ID, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ID, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.ID, "x"},
		{token.COMMA, ","},
		{token.ID, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.ID, "x"},
		{token.PLUS, "+"},
		{token.ID, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ID, "result"},
		{token.ASSIGN, "="},
		{token.ID, "add"},
		{token.LPAREN, "("},
		{token.ID, "five"},
		{token.COMMA, ","},
		{token.ID, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ID, "my_test$123"},
		{token.ASSIGN, "="},
		{token.INT, "15"},
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

func TestOneCharToken(t *testing.T) {
	input := `	!-/*5;
	5 < 10 > 5;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
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

func TestKeywordToken(t *testing.T) {
	input := `if (5 < 10) {
		return true;
	} else {
		return false;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
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

func TestTwoCharToken(t *testing.T) {
	input := `10 == 10;
	10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
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
