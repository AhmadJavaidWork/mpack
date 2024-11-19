package parser

import (
	"fmt"
	"strconv"

	"github.com/ahmadjavaidwork/mpack/ast"
	"github.com/ahmadjavaidwork/mpack/lexer"
	"github.com/ahmadjavaidwork/mpack/token"
)

type Parser struct {
	curToken  token.Token
	peekToken token.Token
	encoding  []string
	l         *lexer.Lexer
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Object {
	o := &ast.Object{Token: p.curToken}

	o.Entries = []*ast.Entry{}

	for !p.curTokenIs(token.RBRACE) {
		p.nextToken()
		entry := p.parseEntry()
		if entry != nil {
			o.Entries = append(o.Entries, entry)
		}

		if entry == nil {
			break
		}
		p.nextToken()
	}

	return o
}

func (p *Parser) parseEntry() *ast.Entry {
	entry := &ast.Entry{
		Token: p.curToken,
		Key:   p.curToken.Literal,
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken()
	entry.Value = p.parseNode()
	return entry
}

func (p *Parser) parseString() *ast.StringLiteral {
	lit := &ast.StringLiteral{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	return lit
}

func (p *Parser) parseNumber() *ast.Number {
	lit := &ast.Number{Token: p.curToken}

	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse '%s' to number", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = val
	return lit
}

func (p *Parser) parseArray() *ast.Array {
	arr := &ast.Array{Token: p.curToken}
	arr.Values = []ast.Node{}

	if p.peekTokenIs(token.RBRACKET) {
		p.nextToken()
		return arr
	}

	p.nextToken()
	arr.Values = append(arr.Values, p.parseNode())

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		arr.Values = append(arr.Values, p.parseNode())
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return arr
}

func (p *Parser) parseNode() ast.Node {
	switch p.curToken.Type {
	case token.STRING:
		return p.parseString()

	case token.NUMBER:
		return p.parseNumber()

	case token.LBRACKET:
		return p.parseArray()

	case token.LBRACE:
		return p.Parse()

	default:
		return nil
	}
}

func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}
	msg := fmt.Sprintf("expected next token to be '%s' got '%s'", tokenType, p.curToken.Type)
	p.errors = append(p.errors, msg)
	return false
}

func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) curTokenIs(tokenType token.TokenType) bool {
	return p.curToken.Type == tokenType
}

func (p *Parser) Errors() []string {
	return p.errors
}