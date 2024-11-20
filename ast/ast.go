package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ahmadjavaidwork/mpack/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Entry struct {
	Token token.Token
	Key   string
	Value Node
}

type Object struct {
	Token   token.Token
	Entries []*Entry
}

func (o *Object) TokenLiteral() string { return o.Token.Literal }
func (o *Object) String() string {
	var out bytes.Buffer

	entries := []string{}
	for _, entry := range o.Entries {
		entries = append(entries, entry.Key+": "+entry.Value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(entries, ", "))
	out.WriteString("}")

	return out.String()
}

type Key struct {
	Token token.Token
	Value string
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *StringLiteral) String() string       { return s.Value }

type IntegerLiteral struct {
	Token token.Token
	Value uint64
	Type  string
}

func (n *IntegerLiteral) TokenLiteral() string { return n.Token.Literal }
func (n *IntegerLiteral) String() string       { return n.Token.Literal }

type Array struct {
	Token  token.Token
	Values []Node
}

func (a *Array) TokenLiteral() string { return a.Token.Literal }
func (a *Array) String() string {
	var out bytes.Buffer

	nodes := []string{}
	for _, node := range a.Values {
		nodes = append(nodes, node.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(nodes, ", "))
	out.WriteString("]")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return fmt.Sprintf("%t", b.Value) }
