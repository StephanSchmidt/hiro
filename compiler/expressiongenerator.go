package compiler

import (
	"strconv"
	"strings"
)

type ExprGenerator struct {
	Sb *strings.Builder
}

func (g *ExprGenerator) visitCall(c *Call) {
	g.Sb.WriteString(c.Name)
	g.Sb.WriteString("(")
	for i, par := range c.Args {
		//Sb.WriteString(par)
		g.visitExpression(par)
		if i < len(c.Args)-1 {
			g.Sb.WriteString(",")
		}
	}
	g.Sb.WriteString(")")
}

func (g *ExprGenerator) visitPrimary(p *Primary) {
	if p.String != nil {
		g.Sb.WriteString(*p.String)
	}
	if p.Float != nil {
		g.Sb.WriteString(strconv.FormatFloat(3.1415, 'E', -1, 64))
	}
	if p.Int != nil {
		g.Sb.WriteString(strconv.Itoa(*p.Int))
	}
	if p.Variable != nil {
		g.Sb.WriteString(*p.Variable)
	}
	if p.SubExpression != nil {
		g.Sb.WriteString("(")
		g.visitExpression(p.SubExpression)
		g.Sb.WriteString(")")
	}
	if p.Call != nil {
		g.visitCall(p.Call)
	}
}
func (g *ExprGenerator) visitUnary(u *Unary) {
	g.Sb.WriteString(u.Op)
	if u.Unary != nil {
		g.visitUnary(u.Unary)
	}
	g.visitPrimary(u.Primary)
}

func (g *ExprGenerator) visitMultiplication(m *Multiplication) {
	if m.Unary != nil {
		g.visitUnary(m.Unary)
	}
	g.Sb.WriteString(m.Op)
	if m.Next != nil {
		g.visitMultiplication(m.Next)
	}
}

func (g *ExprGenerator) visitAddition(a *Addition) {
	g.visitMultiplication(a.Multiplication)
	g.Sb.WriteString(a.Op)
	if a.Next != nil {
		g.visitAddition(a.Next)
	}
}

func (g *ExprGenerator) visitComparison(c *Comparison) {
	g.visitAddition(c.Addition)
	g.Sb.WriteString(c.Op)
	if c.Next != nil {
		g.visitComparison(c.Next)
	}

}
func (g *ExprGenerator) visitEquality(e *Equality) {
	g.visitComparison(e.Comparison)
	g.Sb.WriteString(e.Op)
	if e.Next != nil {
		g.visitEquality(e.Next)
	}
}
func (g *ExprGenerator) visitExpression(e *Expression) {
	if e.Equality != nil {
		g.visitEquality(e.Equality)
	}
}
