package compiler

type Symbols struct {
	Symbols map[string]bool
}

func NewSymbols() *Symbols {
	s := new(Symbols)
	s.Symbols = make(map[string]bool)
	return s
}

func (s *Symbols) add(symbol *string) {
	s.Symbols[*symbol] = false
}

func (s *Symbols) contains(symbol *string) bool {
	_, there := s.Symbols[*symbol]
	return there
}

func (s *Symbols) resolve(symbol *string) {
	s.Symbols[*symbol] = true
}

func (s *Symbols) isResolved(symbol *string) bool {
	if !s.contains(symbol) {
		return false
	}
	return s.Symbols[*symbol]
}
