package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithFreeVarsRight(t *testing.T) {
	e := expr(`a > x`)
	we := ToWithExpression(RightLeftVars(e),
		"call",
		[]string{"a"})

	assert.True(t, we.RightHasFreeVar)
	assert.False(t, we.LeftHasFreeVar)
}

func TestWithFreeVarsLeft(t *testing.T) {
	e := expr(`x > a`)
	we := ToWithExpression(RightLeftVars(e),
		"call",
		[]string{"a"})

	assert.False(t, we.RightHasFreeVar)
	assert.True(t, we.LeftHasFreeVar)
}

func TestWithExpression(t *testing.T) {
	e := expr(`call(a) == b`)
	we := ToWithExpression(RightLeftVars(e),
		"call",
		[]string{"a"})

	assert.Equal(t, "==", we.Op)
	assert.True(t, we.LeftContainsVar)
	assert.False(t, we.RightContainsVar)
	assert.True(t, we.LeftHasVars)
	assert.True(t, we.RightHasVars)
	assert.Equal(t, "a", we.LeftVars[0])
	assert.Equal(t, "b", we.RightVars[0])
	assert.True(t, we.LeftContainsCall)
	assert.False(t, we.RightContainsCall)
}

func Func(f string) *Function {
	parser := participle.MustBuild[Function](participle.UseLookahead(2))
	fu, _ := parser.ParseString("", f)
	return fu
}

func TestUnitTest(t *testing.T) {
	f := Func(
		`fn add(a:int, b:int) -> int:
 		end`)

	e := expr(`add(2,3) > 2`)
	we := ToWithExpression(RightLeftVars(e),
		"add",
		[]string{"a", "b"})
	assert.False(t, we.LeftContainsVar)
	assert.False(t, we.RightContainsVar)
	assert.False(t, we.LeftHasVars)
	assert.False(t, we.RightHasVars)
	assert.True(t, we.LeftContainsCall)
	assert.False(t, we.RightContainsCall)
	state := WithExpressionTypeFor(f, we)
	assert.Equal(t, UnitTest, state)
}

func TestPropTest(t *testing.T) {
	f := Func(
		`fn add(a:int, b:int) -> int:
 		end`)

	e := expr(`add(a,b) > a + b`)
	we := ToWithExpression(RightLeftVars(e),
		"add",
		[]string{"a", "b"})
	assert.True(t, we.LeftContainsVar)
	assert.True(t, we.LeftHasVars)
	assert.True(t, we.LeftContainsCall)
	state := WithExpressionTypeFor(f, we)
	assert.Equal(t, PropTest, state)
}

func TestAssertion(t *testing.T) {
	f := Func(
		`fn add(a:int, b:int) -> int:
 		end`)

	e := expr(`a > 2`)
	we := ToWithExpression(RightLeftVars(e),
		"add",
		[]string{"a", "b"})
	assert.True(t, we.LeftContainsVar)
	assert.False(t, we.RightContainsVar)
	assert.True(t, we.LeftHasVars)
	assert.False(t, we.RightHasVars)
	assert.False(t, we.LeftContainsCall)
	assert.False(t, we.RightContainsCall)
	state := WithExpressionTypeFor(f, we)
	assert.Equal(t, Assertion, state)
}

func TestIllegal(t *testing.T) {
	f := Func(
		`fn add(a:int, b:int) -> int:
 		end`)

	e := expr(`1>2`)
	we := ToWithExpression(RightLeftVars(e),
		"add",
		[]string{"a", "b"})
	state := WithExpressionTypeFor(f, we)
	assert.Equal(t, Illegal, state)
}

func TestIllegalFreeVars(t *testing.T) {
	f := Func(
		`fn add(a:int, b:int) -> int:
 		end`)

	e := expr(`a>x`)
	we := ToWithExpression(RightLeftVars(e),
		"add",
		[]string{"a", "b"})
	state := WithExpressionTypeFor(f, we)
	assert.Equal(t, Illegal, state)
}
