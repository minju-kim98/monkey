package parser

import (
	"fmt"

	"github.com/minju-kim98/monkey/ast"
	"github.com/minju-kim98/monkey/lexer"
	"github.com/minju-kim98/monkey/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	currToken token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // function call
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.ID, p.parseIdentifier)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) registerPrefix(tt token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}

func (p *Parser) registerInfix(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
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
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExpressionStatement()
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

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()

	for !p.currTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
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

	p.peekError(tt)
	return false
}

func (p *Parser) getErrors() []string {
	return p.errors
}

func (p *Parser) peekError(tt token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tt, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
