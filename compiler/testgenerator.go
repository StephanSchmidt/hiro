package compiler

import (
	"github.com/thoas/go-funk"
	"strconv"
	"strings"
)

type TestGenerator struct {
	Sb    *strings.Builder
	Count int
}

func HasUnitTest(ast *HiroAst) bool {
	for _, f := range ast.Functions {
		for _, we := range f.With {
			if we.Parsed.WithType == UnitTest {
				return true
			}
		}
	}
	return false
}

func HasPropTest(ast *HiroAst) bool {
	for _, f := range ast.Functions {
		for _, we := range f.With {
			if we.Parsed.WithType == PropTest {
				return true
			}
		}
	}
	return false
}

func (g *TestGenerator) VisitAst(ast *HiroAst) {
	g.Sb.WriteString(`
import (`)
	if HasUnitTest(ast) {
		g.Sb.WriteString("\"testing\"\n")
	}
	if HasPropTest(ast) {
		g.Sb.WriteString("\"pgregory.net/rapid\"\n")
	}
	g.Sb.WriteString(`)
`)

	for _, fu := range ast.Functions {
		g.visitFunction(fu)
	}
}

func (g *TestGenerator) visitFunction(f *Function) {
	AnnotateFunctionWith(f)
	for _, with := range f.With {
		if with.Parsed.WithType == UnitTest {
			var sb strings.Builder
			var symbols = NewSymbols()

			exprGenerator := &GoGenerator{
				Sb:      &sb,
				Symbols: symbols,
			}
			exprGenerator.visitExpression(with.Expression)

			var esb strings.Builder
			eGenerator := &ExprGenerator{
				Sb: &esb,
			}
			eGenerator.visitExpression(with.Expression)

			g.Count = g.Count + 1
			g.Sb.WriteString("func Test_" + f.Name + "_" + strconv.Itoa(g.Count) + "(t *testing.T) {\n")
			g.Sb.WriteString("if ! (")
			g.Sb.WriteString(sb.String())
			g.Sb.WriteString(`){
	t.Errorf("Assertion failed: ` + esb.String() + `")
}
}`)
			g.Sb.WriteString("\n")
		}
		if with.Parsed.WithType == PropTest {
			var sb strings.Builder
			var symbols = NewSymbols()
			exprGenerator := &GoGenerator{
				Sb:      &sb,
				Symbols: symbols,
			}
			exprGenerator.visitExpression(with.Expression)

			var esb strings.Builder

			eGenerator := &ExprGenerator{
				Sb: &esb,
			}
			eGenerator.visitExpression(with.Expression)

			g.Count = g.Count + 1
			g.Sb.WriteString("func Test_" + f.Name + "_" + strconv.Itoa(g.Count) + "(t *testing.T) {\n")
			g.Sb.WriteString("rapid.Check(t, func(t *rapid.T) {\n")

			vars := funk.UniqString(append(with.Parsed.LeftVars, with.Parsed.RightVars...))
			for _, v := range vars {
				g.Sb.WriteString("")
				typ := f.Parsed.Symbols[v]
				// 		a := rapid.Int().Draw(t, "a").(int)
				g.Sb.WriteString(v + ":= rapid." + strings.Title(typ) + "().Draw(t, \"" + v + "\").(" + typ + ")\n")
			}

			g.Sb.WriteString("if ! (")
			g.Sb.WriteString(sb.String())
			g.Sb.WriteString(`){
t.Fatalf("Test failed: ` + esb.String() + `")
			}
			`)
			g.Sb.WriteString("})\n") // close rapid.Check
			g.Sb.WriteString("}\n")  // close Test_
			g.Sb.WriteString("\n")
		}
	}
}
