package compiler

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProperty(t *testing.T) {
	parser := participle.MustBuild[Expression](participle.UseLookahead(2))
	hiro, _ := parser.ParseString("", "a>0")
	assert.Equal(t, ">", hiro.Equality.Comparison.Op, "Test gemerated")
}
