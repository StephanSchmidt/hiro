package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
)

type GoGenerator struct {
	sb      *strings.Builder
	Symbols *map[string]bool
}

func (g *GoGenerator) visitAst(ast *HiroAst) {
	g.sb.WriteString("package main\n\n")
	g.sb.WriteString("import (\n\t\"fmt\")\n\n")

	for _, fu := range ast.Functions {
		g.visitFunction(fu)
	}
}

func (g *GoGenerator) visitFunction(f *Function) {

	g.sb.WriteString(fmt.Sprintf(`func %s(`, f.Name))
	for i, arg := range f.Args {
		g.sb.WriteString(arg.VarName)
		g.sb.WriteString(" ")
		g.sb.WriteString(arg.VarType)
		if i < len(f.Args)-1 {
			g.sb.WriteString(", ")
		}
	}
	hasReturn := false
	if len(f.Return) != 0 {
		hasReturn = true
	}
	// if f.Body[len(f.Body)-1].Expression != nil {
	//	hasReturn = true
	//}
	if hasReturn {
		g.sb.WriteString(fmt.Sprintf(`) <- chan %s {`, f.Return))
		g.sb.WriteString(fmt.Sprintf("\nres := make(chan %s)\ngo func(){\n defer close(res)\n", f.Return))
		for index, c := range f.Body {
			if index == len(f.Body)-1 {
				g.sb.WriteString("res <- ")
			}
			g.visitCommand(c)
		}
		g.sb.WriteString("}()\n")
		g.sb.WriteString("return res\n")
		g.sb.WriteString("}\n")
	} else {
		g.sb.WriteString(fmt.Sprintf(`) {`))
		g.sb.WriteString("\n")
		for _, c := range f.Body {
			g.visitCommand(c)
		}
		g.sb.WriteString("}\n")
	}
}

func (g *GoGenerator) visitCommand(c *Command) {
	if c.Print != nil {
		v := &VarsChecker{}
		v.visitExpression(c.Print.Expression)
		spew.Dump(v.Vars)
		for _, variable := range v.Vars {
			fmt.Println(variable)
			fmt.Println((*g.Symbols)[variable])
			if (*g.Symbols)[variable] == false {
				(*g.Symbols)[variable] = true
				g.sb.WriteString("var ")
				g.sb.WriteString(variable + "_done")
				g.sb.WriteString(" = <- ")
				g.sb.WriteString(variable)
				g.sb.WriteString("\n")
			}
		}

		g.sb.WriteString("fmt.Println(")
		g.visitExpression(c.Print.Expression)
		// sb.WriteString(c.Print.Expression)
		g.sb.WriteString(")")
		g.sb.WriteString("\n")
	}
	if c.Expression != nil {
		g.visitExpression(c.Expression)
		g.sb.WriteString("\n")
	}
	if c.Call != nil {
		g.visitCall(c.Call)
		g.sb.WriteString("\n")
	}
	if c.Let != nil {
		g.visitLet(c.Let)
		g.sb.WriteString("\n")
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
		g.sb.WriteString(varName + " := make(chan any)\n")
		g.sb.WriteString("go func() {\n")
		g.sb.WriteString("defer close(" + varName + ")\n")
		g.sb.WriteString(varName)
		g.sb.WriteString(" <- ")
		g.visitExpression(l.Expr)
		g.sb.WriteString("}()\n")
		(*g.Symbols)[l.Var] = false
	} else {
		g.sb.WriteString("var ")
		g.sb.WriteString(l.Var + "_done")
		g.sb.WriteString(" = ")
		g.visitExpression(l.Expr)
		(*g.Symbols)[l.Var] = true
	}
}

func (g *GoGenerator) visitCall(c *Call) {
	g.sb.WriteString("(<- ")
	g.sb.WriteString(c.Name)
	g.sb.WriteString("(")
	for i, par := range c.Args {
		//sb.WriteString(par)
		g.visitExpression(par)
		if i < len(c.Args)-1 {
			g.sb.WriteString(",")
		}
	}
	g.sb.WriteString("))")
}

func (g *GoGenerator) visitPrimary(p *Primary) {
	if p.String != nil {
		g.sb.WriteString(*p.String)
	}
	if p.Float != nil {
		g.sb.WriteString(strconv.FormatFloat(3.1415, 'E', -1, 64))
	}
	if p.Int != nil {
		g.sb.WriteString(strconv.Itoa(*p.Int))
	}
	if p.Variable != nil {
		there, _ := (*g.Symbols)[*p.Variable]
		if there {
			g.sb.WriteString(*p.Variable + "_done")
		} else {
			g.sb.WriteString(*p.Variable)
		}
	}
	if p.SubExpression != nil {
		g.sb.WriteString("(")
		g.visitExpression(p.SubExpression)
		g.sb.WriteString(")")
	}
	if p.Call != nil {
		g.visitCall(p.Call)
	}
}
func (g *GoGenerator) visitUnary(u *Unary) {
	g.sb.WriteString(u.Op)
	if u.Unary != nil {
		g.visitUnary(u.Unary)
	}
	g.visitPrimary(u.Primary)
}

func (g *GoGenerator) visitMultiplication(m *Multiplication) {
	if m.Unary != nil {
		g.visitUnary(m.Unary)
	}
	g.sb.WriteString(m.Op)
	if m.Next != nil {
		g.visitMultiplication(m.Next)
	}
}

func (g *GoGenerator) visitAddition(a *Addition) {
	g.visitMultiplication(a.Multiplication)
	g.sb.WriteString(a.Op)
	if a.Next != nil {
		g.visitAddition(a.Next)
	}
}

func (g *GoGenerator) visitComparison(c *Comparison) {
	g.visitAddition(c.Addition)
	g.sb.WriteString(c.Op)
	if c.Next != nil {
		g.visitComparison(c.Next)
	}

}
func (g *GoGenerator) visitEquality(e *Equality) {
	g.visitComparison(e.Comparison)
	g.sb.WriteString(e.Op)
	if e.Next != nil {
		g.visitEquality(e.Next)
	}
}
func (g *GoGenerator) visitExpression(e *Expression) {
	if e.Equality != nil {
		g.visitEquality(e.Equality)
	}
}
