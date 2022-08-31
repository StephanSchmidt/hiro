package compiler

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
)

type GoGenerator struct {
	Sb      *strings.Builder
	Symbols *Symbols
}

func (g *GoGenerator) VisitAst(ast *HiroAst) {
	g.Sb.WriteString("package main\n\n")
	g.Sb.WriteString("import (\n\t\"fmt\"\n)\n\n")

	for _, fu := range ast.Functions {
		g.visitFunction(fu)
	}
}

func (g *GoGenerator) visitFunction(f *Function) {
	AnnotateFunctionWith(f)

	g.Symbols.newScope()
	g.Sb.WriteString(fmt.Sprintf(`func %s(`, f.Name))
	for i, arg := range f.Args {
		g.Sb.WriteString(arg.VarName)
		g.Sb.WriteString(" ")
		g.Sb.WriteString(arg.VarType)
		if i < len(f.Args)-1 {
			g.Sb.WriteString(", ")
		}
	}
	hasReturn := false
	if len(f.Return) > 0 {
		hasReturn = true
	}
	// if f.Body[len(f.Body)-1].Expression != nil {
	//	hasReturn = true
	//}
	if hasReturn {
		g.Sb.WriteString(fmt.Sprintf(`) <-chan %s {`, f.Return))
		g.Sb.WriteString("\n")
		for _, we := range f.With {
			if we.Parsed == nil {

			}
			if we.Parsed != nil && we.Parsed.WithType == Assertion {
				g.Sb.WriteString(` if !(a > 0) {
  panic("Assertion failed: a > 0")
 }`)
			}
		}
		g.Sb.WriteString(fmt.Sprintf("\n res := make(chan %s)\n go func() {\n defer close(res)\n", f.Return))
		for index, c := range f.Body {
			if index == len(f.Body)-1 {
				g.Sb.WriteString(" res <- ")
			}
			g.visitCommand(c)
		}
		g.Sb.WriteString(" }()\n")
		g.Sb.WriteString(" return res\n")
		g.Sb.WriteString("}\n")
	} else {
		g.Sb.WriteString(fmt.Sprintf(`) {`))
		g.Sb.WriteString("\n")
		for _, c := range f.Body {
			g.visitCommand(c)
		}
		g.Sb.WriteString("}\n")
	}
	g.Symbols.backScope()
}

func (g *GoGenerator) visitCommand(c *Command) {
	if c.If != nil {
		g.Sb.WriteString("if ")
		if c.Expression != nil {
			g.visitExpression(c.Expression)
		}
		g.Sb.WriteString("{\n")
		g.Sb.WriteString("}\n")

	}
	if c.Print != nil {
		v := &ExpressionAnalyzer{}
		v.visitExpression(c.Print.Expression)
		spew.Dump(v.Vars)
		for _, variable := range v.Vars {
			if !g.Symbols.isResolved(&variable) {
				g.Sb.WriteString("var ")
				g.Sb.WriteString("_" + variable)
				g.Sb.WriteString(" = <- ")
				g.Sb.WriteString(variable)
				g.Sb.WriteString("\n")
				g.Symbols.resolve(&variable)
			}
		}

		g.Sb.WriteString("fmt.Println(")
		g.visitExpression(c.Print.Expression)
		// Sb.WriteString(c.Print.Expression)
		g.Sb.WriteString(")")
		g.Sb.WriteString("\n")
	}
	if c.Expression != nil {
		g.visitExpression(c.Expression)
		g.Sb.WriteString("\n")
	}
	if c.Call != nil {
		g.visitCall(c.Call)
		g.Sb.WriteString("\n")
	}
	if c.Let != nil {
		g.visitLet(c.Let)
		g.Sb.WriteString("\n")
	}
}

func (g *GoGenerator) visitLet(l *Let) {
	if l.Expr.CheckedForAsync == false {
		ac := &AsyncChecker{
			async: false,
		}
		ac.visitExpression(l.Expr)
		l.Expr.CheckedForAsync = true
		l.Expr.IsAsync = ac.async
	}

	if l.Expr.IsAsync {
		varName := l.Var
		g.Sb.WriteString(varName + " := make(chan any)\n")
		g.Sb.WriteString("go func() {\n")
		g.Sb.WriteString("defer close(" + varName + ")\n")
		g.Sb.WriteString(varName)
		g.Sb.WriteString(" <- ")
		g.visitExpression(l.Expr)
		g.Sb.WriteString("}()\n")
		g.Symbols.add(&l.Var)
	} else {
		g.Sb.WriteString("var ")
		g.Sb.WriteString("_" + l.Var)
		g.Sb.WriteString(" = ")
		g.visitExpression(l.Expr)
		g.Symbols.add(&l.Var)
		g.Symbols.resolve(&l.Var)
	}
}

func (g *GoGenerator) visitCall(c *Call) {
	g.Sb.WriteString("(<- ")
	g.Sb.WriteString(c.Name)
	g.Sb.WriteString("(")
	for i, par := range c.Args {
		//Sb.WriteString(par)
		g.visitExpression(par)
		if i < len(c.Args)-1 {
			g.Sb.WriteString(",")
		}
	}
	g.Sb.WriteString("))")
}

func (g *GoGenerator) visitPrimary(p *Primary) {
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
		if g.Symbols.contains(p.Variable) {
			g.Sb.WriteString("_" + *p.Variable)
		} else {
			g.Sb.WriteString(*p.Variable)
		}
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
func (g *GoGenerator) visitUnary(u *Unary) {
	g.Sb.WriteString(u.Op)
	if u.Unary != nil {
		g.visitUnary(u.Unary)
	}
	g.visitPrimary(u.Primary)
}

func (g *GoGenerator) visitMultiplication(m *Multiplication) {
	if m.Unary != nil {
		g.visitUnary(m.Unary)
	}
	g.Sb.WriteString(m.Op)
	if m.Next != nil {
		g.visitMultiplication(m.Next)
	}
}

func (g *GoGenerator) visitAddition(a *Addition) {
	g.visitMultiplication(a.Multiplication)
	g.Sb.WriteString(a.Op)
	if a.Next != nil {
		g.visitAddition(a.Next)
	}
}

func (g *GoGenerator) visitComparison(c *Comparison) {
	g.visitAddition(c.Addition)
	g.Sb.WriteString(c.Op)
	if c.Next != nil {
		g.visitComparison(c.Next)
	}

}
func (g *GoGenerator) visitEquality(e *Equality) {
	g.visitComparison(e.Comparison)
	g.Sb.WriteString(e.Op)
	if e.Next != nil {
		g.visitEquality(e.Next)
	}
}
func (g *GoGenerator) visitExpression(e *Expression) {
	if e.Equality != nil {
		g.visitEquality(e.Equality)
	}
}
