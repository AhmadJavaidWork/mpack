package lexer

import "github.com/ahmadjavaidwork/mpack/token"

type Lexer struct {
	input        string
	current      byte
	position     int
	readPosition int
}

func newLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readByte()
	return l
}

func (l *Lexer) readByte() {
	if l.readPosition >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpaces()

	switch l.current {
	case 0:
		tok = token.Token{Type: token.EOF, Literal: token.EOF}
	case '"':
		l.readByte()
		tok = token.Token{Type: token.STRING, Literal: l.readUntil('"')}
	case ':':
		tok = token.Token{Type: token.COLON, Literal: ":"}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: ","}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: "}"}
	case '[':
		tok = token.Token{Type: token.LBRACKET, Literal: "["}
	case ']':
		tok = token.Token{Type: token.RBRACKET, Literal: "]"}
	default:
		if isDigit(l.current) {
			tok = token.Token{Type: token.NUMBER, Literal: l.readUntil(',')}
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.current)}
		}
	}

	l.readByte()

	return tok
}

func (l *Lexer) skipWhiteSpaces() {
	for l.current == '\n' || l.current == '\t' || l.current == '\r' || l.current == ' ' {
		l.readByte()
	}
}

func (l *Lexer) readUntil(end byte) string {
	pos := l.position
	for l.current != end {
		l.readByte()
	}
	return l.input[pos:l.position]
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
