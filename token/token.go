package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"

	// Identifiers
	ID  = "ID"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	SEMICOLON = ";"
	COMMA     = ","

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Essential
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)
