package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Expr(e string) *Expression {
	parser := participle.MustBuild[Expression](participle.UseLookahead(2))
	ex, _ := parser.ParseString("", e)
	return ex
}

func TestVarChecker(t *testing.T) {
	v := &ExpressionAnalyzer{}
	v.visitExpression(Expr(`a > 0`))
	assert.True(t, len(v.Vars) == 1)
	assert.Equal(t, "a", v.Vars[0])
}

func TestVarCheckerWithCall(t *testing.T) {
	v := &ExpressionAnalyzer{}
	v.visitExpression(Expr(`call(a) > 0`))
	assert.NotNilf(t, v.Vars, "Var in call found")
	assert.True(t, len(v.Vars) == 1)
	assert.Equal(t, "a", v.Vars[0])
}

func TestVarCheckerForCall(t *testing.T) {
	v := &ExpressionAnalyzer{}
	v.visitExpression(Expr(`call(a) > 0`))
	assert.NotNilf(t, v.Calls, "Call found")
	assert.True(t, len(v.Calls) == 1)
	assert.Equal(t, "call", v.Calls[0])
}
