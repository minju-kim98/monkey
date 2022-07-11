package parser

import (
	"github.com/minju-kim98/monkey/ast"
	"github.com/minju-kim98/monkey/lexer"
	"github.com/minju-kim98/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Root {
	return nil
}
