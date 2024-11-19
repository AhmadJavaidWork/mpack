package lexer

import (
	"testing"

	"github.com/ahmadjavaidwork/mpack/token"
)

func TestLexer(t *testing.T) {
	input := `{
					"userName": "Martin",
					"favoriteNumber": 1337,
					"interests": [
						"daydreaming",
						"hacking"
					]
				}
			`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "userName",
		},
		{
			expectedType:    token.COLON,
			expectedLiteral: ":",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "Martin",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "favoriteNumber",
		},
		{
			expectedType:    token.COLON,
			expectedLiteral: ":",
		},
		{
			expectedType:    token.NUMBER,
			expectedLiteral: "1337",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "interests",
		},
		{
			expectedType:    token.COLON,
			expectedLiteral: ":",
		},
		{
			expectedType:    token.LBRACKET,
			expectedLiteral: "[",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "daydreaming",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "hacking",
		},
		{
			expectedType:    token.RBRACKET,
			expectedLiteral: "]",
		},
		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		},
	}

	l := newLexer(input)
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Type)
		}
	}
}
