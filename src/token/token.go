package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifier
	IDENT = "IDENT" // add, foobar, x, y
	INT = "INT"     // 123456

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)