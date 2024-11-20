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
	TRUE   = "TRUE"
	FALSE  = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdentifier(keyword string) TokenType {
	if t, ok := keywords[keyword]; ok {
		return t
	}
	return ILLEGAL
}
