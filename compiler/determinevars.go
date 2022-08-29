package compiler

type ExpressionAnalyzer struct {
	Vars  []string
	Calls []string
}

func (v *ExpressionAnalyzer) visitExpression(e *Expression) {
	if e.Equality != nil {
		v.visitEquality(e.Equality)
	}
}

func (v *ExpressionAnalyzer) visitEquality(e *Equality) {
	if e.Comparison != nil {
		v.visitComparison(e.Comparison)
	}
	if e.Next != nil {
		v.visitEquality(e.Next)
	}
}

func (v *ExpressionAnalyzer) visitPrimary(p *Primary) {
	if p.Variable != nil {
		v.Vars = append(v.Vars, *p.Variable)
	}
	if p.SubExpression != nil {
		v.visitExpression(p.SubExpression)
	}
	if p.Call != nil {
		v.Calls = append(v.Calls, p.Call.Name)
		for _, arg := range p.Call.Args {
			v.visitExpression(arg)
		}
	}
}

func (v *ExpressionAnalyzer) visitComparison(c *Comparison) {
	if c.Addition != nil {
		v.visitAddition(c.Addition)
	}
	if c.Next != nil {
		v.visitComparison(c.Next)
	}
}

func (v *ExpressionAnalyzer) visitAddition(a *Addition) {
	if a.Multiplication != nil {
		v.visitMultiplication(a.Multiplication)
	}
	if a.Next != nil {
		v.visitAddition(a.Next)
	}
}

func (v *ExpressionAnalyzer) visitMultiplication(m *Multiplication) {
	if m.Unary != nil {
		v.visitUnary(m.Unary)
	}
	if m.Next != nil {
		v.visitMultiplication(m.Next)
	}
}

func (v *ExpressionAnalyzer) visitUnary(u *Unary) {
	if u.Unary != nil {
		v.visitUnary(u.Unary)
	}
	if u.Primary != nil {
		v.visitPrimary(u.Primary)
	}
}
