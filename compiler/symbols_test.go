package compiler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymbols(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.add(&sym)
	fmt.Println(s.contains(&sym))
	assert.True(t, s.contains(&sym), "Symbols should contain symbol after added symbol")
}
