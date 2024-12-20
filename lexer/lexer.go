package lexer

import (
	"github.com/ahmadjavaidwork/mpack/token"
)

type Lexer struct {
	input        string
	current      byte
	position     int
	readPosition int
}

func New(input string) *Lexer {
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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpaces()

	switch l.current {
	case 0:
		tok = token.Token{Type: token.EOF, Literal: token.EOF}
	case '"':
		termintors := map[byte]bool{
			'"': true,
		}
		l.readByte()
		tok = token.Token{Type: token.STRING, Literal: l.readUntil(termintors)}
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
		if isLetter(l.current) {
			termintors := map[byte]bool{
				',':  true,
				'}':  true,
				'\n': true,
			}
			tok.Literal = l.readUntil(termintors)
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.current) || l.current == '-' {
			termintors := map[byte]bool{
				',':  true,
				'}':  true,
				'\n': true,
			}
			tok = token.Token{Type: token.NUMBER, Literal: l.readUntil(termintors)}
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

func (l *Lexer) readUntil(termintors map[byte]bool) string {
	pos := l.position
	for _, ok := termintors[l.current]; !ok; _, ok = termintors[l.current] {
		l.readByte()
	}
	return l.input[pos:l.position]
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}
