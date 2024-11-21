package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/ahmadjavaidwork/mpack/ast"
	"github.com/ahmadjavaidwork/mpack/lexer"
	"github.com/ahmadjavaidwork/mpack/parser"
)

func Encode(json string) []byte {
	l := lexer.New(string(json))
	p := parser.NewParser(l)
	obj := p.Parse()

	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	return encodeObj(*obj)
}

func encodeObj(obj ast.Object) []byte {
	res := []byte{}

	entriesLength := len(obj.Entries)
	if entriesLength < 16 {
		res = append(res, 128+uint8(entriesLength))
	} else if entriesLength < int(math.Pow(2, 16)) {
		res = append(res, 0xde)
		res = binary.BigEndian.AppendUint16(res, uint16(entriesLength))
	} else {
		res = append(res, 0xdf)
		res = binary.BigEndian.AppendUint32(res, uint32(entriesLength))
	}

	for _, entry := range obj.Entries {
		res = append(res, encodeString(entry.Key)...)
		res = append(res, encodeEntry(entry.Value)...)
		break
	}

	return res
}

func encodeEntry(node ast.Node) []byte {
	switch v := node.(type) {
	case *ast.StringLiteral:
		return encodeString(v.Value)
	default:
		return []byte{}
	}
}

func encodeString(str string) []byte {
	res := []byte{}
	strLen := len(str)
	if strLen < 32 {
		res = append(res, 160+uint8(strLen))
	} else if strLen < int(math.Pow(2, 8)) {
		res = append(res, 0xd9)
		res = append(res, byte(strLen))
	} else if strLen < int(math.Pow(2, 16)) {
		res = append(res, 0xda)
		res = binary.BigEndian.AppendUint16(res, uint16(strLen))
	} else {
		res = append(res, 0xdb)
		res = binary.BigEndian.AppendUint32(res, uint32(strLen))
	}

	res = append(res, []byte(str)...)
	return res
}
