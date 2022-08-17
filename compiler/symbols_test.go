package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymbols(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.add(&sym)
	assert.True(t, s.contains(&sym), "Symbols should contain symbol after added symbol")
}

func TestAsyncSymbols(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.add(&sym)
	s.resolve(&sym)
	assert.True(t, s.isResolved(&sym), "Symbols should contain symbol after added symbol")
}
