package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahmadjavaidwork/mpack/ast"
	"github.com/ahmadjavaidwork/mpack/lexer"
	"github.com/ahmadjavaidwork/mpack/token"
)

type Parser struct {
	curToken  token.Token
	peekToken token.Token
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

func (p *Parser) parseNode() ast.Node {
	switch p.curToken.Type {
	case token.STRING:
		return p.parseStringLiteral()

	case token.NUMBER:
		return p.parseIntegerLiteral()

	case token.TRUE, token.FALSE:
		return p.parseBoolean()

	case token.LBRACKET:
		return p.parseArray()

	case token.LBRACE:
		return p.Parse()

	default:
		msg := fmt.Sprintf("illegal token: %s", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
}

func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	lit := &ast.StringLiteral{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	return lit
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	num := p.curToken.Literal
	if strings.Contains(num, "-") {
		lit.Value, lit.Type = p.parseSignedInt(num)
	} else {
		lit.Value, lit.Type = p.parseUnsignedInt(num)
	}

	if lit.Value == 0 {
		return nil
	}

	return lit
}

func (p *Parser) parseSignedInt(num string) (uint64, string) {
	val, err := strconv.ParseInt(num, 0, 8)
	if err == nil {
		return uint64(val), "int8"
	}

	val, err = strconv.ParseInt(num, 0, 16)
	if err == nil {
		return uint64(val), "int16"
	}

	val, err = strconv.ParseInt(num, 0, 32)
	if err == nil {
		return uint64(val), "int32"
	}

	val, err = strconv.ParseInt(num, 0, 64)
	if err == nil {
		return uint64(val), "int64"
	}

	msg := fmt.Sprintf("could not parse %s as int64", num)
	p.errors = append(p.errors, msg)

	return 0, ""
}

func (p *Parser) parseUnsignedInt(num string) (uint64, string) {
	val, err := strconv.ParseUint(num, 0, 8)
	if err == nil {
		return val, "uint8"
	}

	val, err = strconv.ParseUint(num, 0, 16)
	if err == nil {
		return val, "uint16"
	}

	val, err = strconv.ParseUint(num, 0, 32)
	if err == nil {
		return val, "uint32"
	}

	val, err = strconv.ParseUint(num, 0, 8)
	if err == nil {
		return val, "uint64"
	}

	msg := fmt.Sprintf("could not parse %s as uint64", num)
	p.errors = append(p.errors, msg)

	return 0, ""
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

func (p *Parser) parseBoolean() *ast.Boolean {
	return &ast.Boolean{Token: p.curToken, Value: isTrue(p.curToken.Type)}
}

func isTrue(t token.TokenType) bool {
	return t == token.TRUE
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
