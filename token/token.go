package token

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	COLON = ":"
	COMMA = ","

	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	STRING = "STRING"
	NUMBER = "NUMBER"
	ARRAY  = "ARRAY"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
