package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"go/format"
	"strings"
	"testing"
)

func Generate(source string) string {
	parser := participle.MustBuild[HiroAst](participle.UseLookahead(2))
	code, _ := parser.ParseString("", source)
	var sb strings.Builder
	var symbols = NewSymbols()
	goGenerator := &GoGenerator{
		Sb:      &sb,
		Symbols: symbols,
	}
	goGenerator.VisitAst(code)
	return sb.String()
}

func TestGenerate(t *testing.T) {
	source := `
	fn add(a:int, b:int) -> int:
		a + b
	end
	`
	expected := `
package main

import (
	"fmt"
)

func add(a int, b int) <-chan int {

	res := make(chan int)
    go func() {
      defer close(res)
      res <- a+b
    }()
    return res
}
`
	assert.Equal(t, formatSource(expected), formatSource(Generate(source)))
}

func TestFunctionParse(t *testing.T) {
	f := Func(`
	fn add(a:int, b:int) with
		a > 0
		-> int:
	end
	`)
	assert.Equal(t, "add", f.Name)
	assert.Equal(t, "int", f.Return)
	assert.Equal(t, 1, len(f.With))
}

func formatSource(source string) string {
	formatted, _ := format.Source([]byte(source))
	return string(formatted)
}

func TestGenerateAssertion(t *testing.T) {
	source := `
	fn add(a:int, b:int) with
		b > 0
		-> int:
		a + b
	end
	`
	expected := `package main
import ( "fmt" )
func add(a int, b int) <-chan int {
	if !(b > 0) {	
		panic("Assertion failed: b>0")	
	}
    res := make(chan int)
    go func() {
       defer close(res)
       res <- a + b
    }()
    return res
}`
	assert.Equal(t, formatSource(expected), formatSource(Generate(source)))
}
