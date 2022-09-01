package compiler

import (
	"strconv"
	"strings"
)

type TestGenerator struct {
	Sb    *strings.Builder
	Count int
}

func (g *TestGenerator) VisitAst(ast *HiroAst) {
	g.Sb.WriteString("package main\n\n")
	g.Sb.WriteString("import (\n\t\"testing\"\n)\n\n")

	for _, fu := range ast.Functions {
		g.visitFunction(fu)
	}
}

func (g *TestGenerator) visitFunction(f *Function) {
	AnnotateFunctionWith(f)
	for _, with := range f.With {
		if with.Parsed.WithType == UnitTest {
			var sb strings.Builder
			exprGenerator := &ExprGenerator{
				Sb: &sb,
			}
			exprGenerator.visitExpression(with.Expression)

			g.Count = g.Count + 1
			g.Sb.WriteString("func Test_" + f.Name + "_" + strconv.Itoa(g.Count) + "(t *testing.T) {\n")
			g.Sb.WriteString("if ! ")
			g.Sb.WriteString(sb.String())
			g.Sb.WriteString(`{
	t.Errorf("Assertion failed: ` + sb.String() + `")
}
}`)
			g.Sb.WriteString("\n")
		}
	}
}
