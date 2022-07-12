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

func (p *Parser) ParseProgram() *ast.Root {
	root := &ast.Root{}
	root.Statements = []ast.Statement{}

	for !p.currTokenTypeIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}
		p.nextToken()
	}

	return root
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStmt()
	default:
		return nil
	}
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectedPeekType(token.ID) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectedPeekType(token.ASSIGN) {
		return nil
	}

	for !p.currTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currTokenTypeIs(tt token.TokenType) bool {
	return p.currToken.Type == tt
}

func (p *Parser) peekTokenTypeIs(tt token.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) expectedPeekType(tt token.TokenType) bool {
	if p.peekTokenTypeIs(tt) {
		p.nextToken()
		return true
	}
	return false
}
