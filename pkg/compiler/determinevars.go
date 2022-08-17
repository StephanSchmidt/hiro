package pkg

type VarsChecker struct {
	Vars []string
}

func (v *VarsChecker) visitExpression(e *Expression) {
	if e.Equality != nil {
		v.visitEquality(e.Equality)
	}
}

func (v *VarsChecker) visitEquality(e *Equality) {
	if e.Comparison != nil {
		v.visitComparison(e.Comparison)
	}
	if e.Next != nil {
		v.visitEquality(e.Next)
	}
}

func (v *VarsChecker) visitPrimary(p *Primary) {
	if p.Variable != nil {
		v.Vars = append(v.Vars, *p.Variable)
	}
}

func (v *VarsChecker) visitComparison(c *Comparison) {
	if c.Addition != nil {
		v.visitAddition(c.Addition)
	}
	if c.Next != nil {
		v.visitComparison(c.Next)
	}
}

func (v *VarsChecker) visitAddition(a *Addition) {
	if a.Multiplication != nil {
		v.visitMultiplication(a.Multiplication)
	}
	if a.Next != nil {
		v.visitAddition(a.Next)
	}
}

func (v *VarsChecker) visitMultiplication(m *Multiplication) {
	if m.Unary != nil {
		v.vititUnary(m.Unary)
	}
	if m.Next != nil {
		v.visitMultiplication(m.Next)
	}
}

func (v *VarsChecker) vititUnary(u *Unary) {
	if u.Unary != nil {
		v.vititUnary(u.Unary)
	}
	if u.Primary != nil {
		v.visitPrimary(u.Primary)
	}
}
