package compiler

type Symbols struct {
	Index   int
	Symbols []map[string]bool
}

func NewSymbols() *Symbols {
	s := new(Symbols)
	s.Symbols = append(s.Symbols, make(map[string]bool))
	return s
}

func (s *Symbols) newScope() {
	s.Index = s.Index + 1
	s.Symbols = append(s.Symbols, make(map[string]bool))
}

func (s *Symbols) backScope() {
	if s.Index > 0 {
		s.Index = s.Index - 1
		s.Symbols = s.Symbols[:len(s.Symbols)-1]

	}
}

func (s *Symbols) add(symbol *string) {
	s.Symbols[s.Index][*symbol] = false
}

func (s *Symbols) contains(symbol *string) bool {
	var i = s.Index
	for i >= 0 {
		_, there := s.Symbols[i][*symbol]
		if there {
			return true
		}
		i = i - 1
	}
	return false
}

func (s *Symbols) resolve(symbol *string) {
	s.Symbols[s.Index][*symbol] = true
}

func (s *Symbols) isResolved(symbol *string) bool {
	if !s.contains(symbol) {
		return false
	}
	return s.Symbols[s.Index][*symbol]
}
