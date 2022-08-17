package pkg

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}

func TestSymbols(t *testing.T) {
	s := NewSymbols()
	sym := "a"
	s.add(&sym)
	assert.OK(t, s.contains(&sym) == true, "Symbols should contain symbol after adding")
}
