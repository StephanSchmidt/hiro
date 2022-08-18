package compiler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymbols(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	symb := "b"
	s.add(&sym)
	assert.True(t, s.contains(&sym), "Symbols should contain symbol after added symbol")
	assert.False(t, s.contains(&symb), "Symbols should not contain symbol without adding")
}

func TestAsyncSymbols(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.add(&sym)
	assert.False(t, s.isResolved(&sym), "Symbols should contain unresolved symbol after added symbol")
	s.resolve(&sym)
	assert.True(t, s.isResolved(&sym), "Symbols should contain resolved symbol after being resolved")
}

func TestScope(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.add(&sym)
	s.newScope()
	assert.True(t, s.contains(&sym), "Symbols should contain symbol after added symbol")
}

func TestBackScope(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.newScope()
	s.add(&sym)
	s.backScope()
	assert.False(t, s.contains(&sym), "Symbols should contain symbol after added symbol")
}
