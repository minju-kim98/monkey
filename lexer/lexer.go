package lexer

import "github.com/minju-kim98/monkey/token"

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextChar()
	return l
}

func (l *Lexer) nextChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Type = token.EOF
		tok.Literal = ""

	default:
		if isLetter(l.ch) {
			tok.Literal = l.checkWord()
			tok.Type = token.CheckWordType(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.nextChar()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextChar()
	}
}

func (l *Lexer) checkWord() string {
	position := l.position
	for isLetter(l.ch) {
		l.nextChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}
