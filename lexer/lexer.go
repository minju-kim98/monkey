package lexer

import (
	"github.com/minju-kim98/monkey/token"
)

type Lexer struct {
	input        string
	position     int
	peekPosition int
	ch           byte
}

// nextChar() moves current position to next position
// readChar() in textbook
func (l *Lexer) nextChar() {
	if l.peekPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.peekPosition]
	}
	l.position = l.peekPosition
	l.peekPosition++
}

// getNextChar() gets value of next char
// peakChar() in textbook
func (l *Lexer) getNextChar() byte {
	if l.peekPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.peekPosition]
	}
}

// NextToken() changes input to token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if twoChar, literal := l.twoCharToken(l.ch); twoChar {
			tok = newTokenString(token.EQ, literal)
		} else {
			tok = newTokenString(token.ASSIGN, literal)
		}
	case '+':
		tok = newTokenChar(token.PLUS, l.ch)
	case '-':
		tok = newTokenChar(token.MINUS, l.ch)
	case '*':
		tok = newTokenChar(token.ASTERISK, l.ch)
	case '/':
		tok = newTokenChar(token.SLASH, l.ch)
	case '!':
		if twoChar, literal := l.twoCharToken(l.ch); twoChar {
			tok = newTokenString(token.NEQ, literal)
		} else {
			tok = newTokenString(token.BANG, literal)
		}
	case '<':
		tok = newTokenChar(token.LT, l.ch)
	case '>':
		tok = newTokenChar(token.GT, l.ch)
	case ';':
		tok = newTokenChar(token.SEMICOLON, l.ch)
	case ',':
		tok = newTokenChar(token.COMMA, l.ch)
	case '(':
		tok = newTokenChar(token.LPAREN, l.ch)
	case ')':
		tok = newTokenChar(token.RPAREN, l.ch)
	case '{':
		tok = newTokenChar(token.LBRACE, l.ch)
	case '}':
		tok = newTokenChar(token.RBRACE, l.ch)

	case 0:
		tok.Type = token.EOF
		tok.Literal = ""

	default:
		if isLetter(l.ch) {
			tok.Literal = l.checkWord()
			tok.Type = token.CheckWordType(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok = newTokenString(token.INT, l.getNum())
			return tok
		} else {
			tok = newTokenChar(token.ILLEGAL, l.ch)
		}
	}

	l.nextChar()

	return tok
}

// skipWhiteSpace() helps get through white spaces
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.nextChar()
	}
}

// checkWord() reads entire identifier/keyword name
// readIdentifier() in textbook
func (l *Lexer) checkWord() string {
	position := l.position
	for isLetter(l.ch) {
		l.nextChar()
	}
	return l.input[position:l.position]
}

// getNum() reads entire number
func (l *Lexer) getNum() string {
	position := l.position
	for isDigit(l.ch) {
		l.nextChar()
	}
	return l.input[position:l.position]
}

// twoCharToken() helps proceed token with two characters
func (l *Lexer) twoCharToken(ch byte) (bool, string) {
	if l.getNextChar() == '=' {
		l.nextChar()
		return true, string(ch) + string(l.ch)
	}
	return false, string(ch)
}

// New() initializes the Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextChar()
	return l
}

// isLetter() checks if the input char is an letter
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// isLetter() checks if the input char is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newTokenChar() helps initialize new token which literal is char
func newTokenChar(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}

// newTokenChar() helps initialize new token which literal is string
func newTokenString(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
