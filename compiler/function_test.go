package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	parser := participle.MustBuild[Function](participle.UseLookahead(2))
	f, _ := parser.ParseString("", `
fn Double(a:int) -> int:
	a + a
end
`)
	assert.NotNil(t, f)
	assert.Equal(t, "Double", f.Name)
}

func TestReturn(t *testing.T) {
	parser := participle.MustBuild[Function](participle.UseLookahead(2))
	f, _ := parser.ParseString("", `
fn Double(a:int) -> int:
	a + a
end
`)
	assert.NotNil(t, f)
	assert.Equal(t, "int", f.Return)
}

func TestWith(t *testing.T) {
	parser := participle.MustBuild[Function](participle.UseLookahead(2))
	f, _ := parser.ParseString("", `
fn Double(a:int) 
	with a>0 
	-> int:
	a + a
end
`)
	assert.NotNil(t, f)
	assert.NotNil(t, f.With)
	assert.Equal(t, 1, len(f.With))
	assert.Equal(t, ">", f.With[0].Equality.Comparison.Op)
	assert.Equal(t, "a", *f.With[0].Equality.Comparison.Addition.Multiplication.Unary.Primary.Variable)
	assert.Equal(t, 0, *f.With[0].Equality.Comparison.Next.Addition.Multiplication.Unary.Primary.Int)
}

func TestTwoWith(t *testing.T) {
	parser := participle.MustBuild[Function](participle.UseLookahead(2))
	f, _ := parser.ParseString("", `
fn Double(a:int) with 
	a>0,
	a<100
	-> int:
	a + a
end
`)
	assert.NotNil(t, f)
	assert.NotNil(t, f.With)
	assert.Equal(t, 2, len(f.With))
	assert.Equal(t, ">", f.With[0].Equality.Comparison.Op)
	assert.Equal(t, "a", *f.With[0].Equality.Comparison.Addition.Multiplication.Unary.Primary.Variable)
	assert.Equal(t, 0, *f.With[0].Equality.Comparison.Next.Addition.Multiplication.Unary.Primary.Int)

	assert.Equal(t, "<", f.With[1].Equality.Comparison.Op)
	assert.Equal(t, "a", *f.With[1].Equality.Comparison.Addition.Multiplication.Unary.Primary.Variable)
	assert.Equal(t, 100, *f.With[1].Equality.Comparison.Next.Addition.Multiplication.Unary.Primary.Int)

}
