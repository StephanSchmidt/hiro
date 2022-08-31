package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func expr(e string) *Expression {
	parser := participle.MustBuild[Expression](participle.UseLookahead(2))
	ex, _ := parser.ParseString("", e)
	return ex
}

func TestTwoWithIsEqualityOrComparisonForComparison(t *testing.T) {
	assert.True(t, IsEqualityOrComparison(expr(`a > 0`)))
}

func TestTwoWithIsEqualityOrComparisonForEquality(t *testing.T) {
	assert.True(t, IsEqualityOrComparison(expr(`a == 0`)))
}

func TestTwoWithIsEqualityOrComparisonForAddition(t *testing.T) {
	assert.False(t, IsEqualityOrComparison(expr(`a + 1`)))
}

func TestTwoWithIsEqualityOrComparisonWithAddition(t *testing.T) {
	assert.True(t, IsEqualityOrComparison(expr(`a + 1 > 1 + 1`)))
}

func TestRightLeftComparison(t *testing.T) {
	e := expr(`a > b`)
	rl := RightLeftVars(e)
	assert.Equal(t, ">", rl.Op)
	assert.Equal(t, "a", rl.LeftVars[0])
	assert.Equal(t, "b", rl.RightVars[0])
}

func TestRightLeftEquality(t *testing.T) {
	e := expr(`a == b`)
	rl := RightLeftVars(e)
	assert.Equal(t, "==", rl.Op)
	assert.Equal(t, "a", rl.LeftVars[0])
	assert.Equal(t, "b", rl.RightVars[0])
}

func TestAnnotateWith(t *testing.T) {
	f := Func(`
	fn add(a:int, b:int) with
		a > 0,
        add(a,b) == a + b,
        add(2,3) == 5
		-> int:
		a + b
	end
	`)
	AnnotateFunctionWith(f)
	assert.NotNil(t, f.With[0].Parsed)
	assert.Equal(t, Assertion, f.With[0].Parsed.WithType)
	assert.Equal(t, PropTest, f.With[1].Parsed.WithType)
	assert.Equal(t, UnitTest, f.With[2].Parsed.WithType)
}
