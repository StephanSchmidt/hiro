package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpressionBiggerThanWithVar(t *testing.T) {
	parser := participle.MustBuild[Expression](participle.UseLookahead(2))
	hiro, _ := parser.ParseString("", `a>0`)
	assert.Equal(t, "a", *hiro.Equality.Comparison.Addition.Multiplication.Unary.Primary.Variable)
	assert.Equal(t, ">", hiro.Equality.Comparison.Op, "Test gemerated")
	assert.Equal(t, 0, *hiro.Equality.Comparison.Next.Addition.Multiplication.Unary.Primary.Int)
}

func TestExpressionComparisonWithVar(t *testing.T) {
	parser := participle.MustBuild[Expression](participle.UseLookahead(2))
	hiro, _ := parser.ParseString("", `a == 1`)
	assert.Equal(t, "a", *hiro.Equality.Comparison.Addition.Multiplication.Unary.Primary.Variable)
	assert.Equal(t, "==", hiro.Equality.Op)
	assert.Equal(t, 1, *hiro.Equality.Next.Comparison.Addition.Multiplication.Unary.Primary.Int)
}

func TestExpressionComparisonWithVarWithUnary(t *testing.T) {
	parser := participle.MustBuild[Expression](participle.UseLookahead(2))
	hiro, _ := parser.ParseString("", `a == -1`)
	assert.Equal(t, "a", *hiro.Equality.Comparison.Addition.Multiplication.Unary.Primary.Variable)
	assert.Equal(t, "==", hiro.Equality.Op)
	assert.Equal(t, "-", hiro.Equality.Next.Comparison.Addition.Multiplication.Unary.Op)
	assert.Equal(t, 1, *hiro.Equality.Next.Comparison.Addition.Multiplication.Unary.Unary.Primary.Int)
}
