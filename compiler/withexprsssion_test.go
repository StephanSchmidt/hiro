package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithExpression(t *testing.T) {
	e := expr(`call(a) == b`)
	we := ToWithExpression(RightLeftVars(e), "call")

	assert.Equal(t, "==", we.Op)
	assert.Equal(t, "a", we.LeftVars[0])
	assert.Equal(t, "b", we.RightVars[0])
	assert.True(t, we.LeftContainsCall)
	assert.False(t, we.RightContainsCall)
}
