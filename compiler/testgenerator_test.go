package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func GenerateUnitTests(source string) string {
	parser := participle.MustBuild[HiroAst](participle.UseLookahead(2))
	code, _ := parser.ParseString("", source)
	var sb strings.Builder

	goGenerator := &TestGenerator{
		Sb:    &sb,
		Count: 0,
	}
	goGenerator.VisitAst(code)
	return sb.String()
}

func TestGenerateUnitTest(t *testing.T) {
	source := `
	fn add(a:int, b:int) with
		add(2,3) == 5
		-> int:
		a + b
	end
	`
	expected := `package main
import ( "testing" )
func Test_add_1(t *testing.T) {
  if ! add(2,3)==5 {
  	t.Errorf("Assertion failed: add(2,3)==5")
  }
}`
	// 	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	assert.Equal(t, formatSource(expected), formatSource(GenerateUnitTests(source)))
}
