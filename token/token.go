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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	LT = "<"
	GT = ">"

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func CheckWordType(word string) TokenType {
	if tok, kw := keywords[word]; kw {
		return tok
	}
	return ID
}
