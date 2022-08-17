package pkg

type AsyncChecker struct {
	async bool
}

func (g *AsyncChecker) visitExpression(e *Expression) {
	if e.Equality != nil {
		g.visitEquality(e.Equality)
	}
}

func (g *AsyncChecker) visitEquality(e *Equality) {
	if e.Comparison != nil {
		g.visitComparison(e.Comparison)
	}
	if e.Next != nil {
		g.visitEquality(e.Next)
	}
}

func (g *AsyncChecker) visitPrimary(p *Primary) {
	if p.Call != nil {
		g.async = true
	}
}

func (g *AsyncChecker) visitComparison(c *Comparison) {
	if c.Addition != nil {
		g.visitAddition(c.Addition)
	}
	if c.Next != nil {
		g.visitComparison(c.Next)
	}
}

func (g *AsyncChecker) visitAddition(a *Addition) {
	if a.Multiplication != nil {
		g.visitMultiplication(a.Multiplication)
	}
	if a.Next != nil {
		g.visitAddition(a.Next)
	}
}

func (g *AsyncChecker) visitMultiplication(m *Multiplication) {
	if m.Unary != nil {
		g.vititUnary(m.Unary)
	}
	if m.Next != nil {
		g.visitMultiplication(m.Next)
	}
}

func (g *AsyncChecker) vititUnary(u *Unary) {
	if u.Unary != nil {
		g.vititUnary(u.Unary)
	}
	if u.Primary != nil {
		g.visitPrimary(u.Primary)
	}
}
