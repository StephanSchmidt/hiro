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
	assert.Equal(t, "a", rl.LeftVars[0])
	assert.Equal(t, "b", rl.RightVars[0])
}

func TestRightLeftEquality(t *testing.T) {
	e := expr(`a == b`)
	rl := RightLeftVars(e)
	assert.Equal(t, "a", rl.LeftVars[0])
	assert.Equal(t, "b", rl.RightVars[0])
}
