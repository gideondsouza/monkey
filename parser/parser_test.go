package parser

import (
	"testing"

	"github.com/gideondsouza/monkey/lexer"
)

func TestLetStatments(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383	
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

}
