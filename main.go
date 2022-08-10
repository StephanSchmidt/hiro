package main

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/davecgh/go-spew/spew"
	"os"
	"strconv"
	"strings"
)

func visitCommand(sb *strings.Builder, c *Command) {
	if c.Print != nil {
		sb.WriteString("fmt.Println(")
		visitExpression(sb, c.Print.Expression)
		sb.WriteString(")\n")
	}
	if c.Expression != nil {
		visitExpression(sb, c.Expression)
	}
	if c.Call != nil {
		visitCall(sb, c.Call)
	}
}

func visitCall(sb *strings.Builder, c *Call) {
	sb.WriteString(c.Name)
	sb.WriteString("(")
	for i, par := range c.Args {
		visitExpression(sb, par)
		if i < len(c.Args)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")\n")
}

func visitOpCmp(sb *strings.Builder, o *OpCmp) {
	sb.WriteString(string(o.Operator))
	visitCmp(sb, o.Cmp)
}

func visitCmp(sb *strings.Builder, c *Cmp) {
	if c.Left != nil {
		visitTerm(sb, c.Left)
	}
	if c.Right != nil {
		for _, ot := range c.Right {
			visitOpTerm(sb, ot)
		}
	}
}

func visitValue(sb *strings.Builder, v *Value) {

	if v.Variable != nil {
		sb.WriteString(*v.Variable)
	}
	if v.String != nil {
		sb.WriteString(*v.String)
	}
	if v.Number != nil {
		sb.WriteString(strconv.Itoa(*v.Number))
	}
	if v.Subexpression != nil {
		sb.WriteString("(")
		visitExpression(sb, v.Subexpression)
		sb.WriteString(")")
	}
}

func visitFactor(sb *strings.Builder, f *Factor) {
	if f.Base != nil {
		visitValue(sb, f.Base)
	}
	if f.Exponent != nil {
		sb.WriteString("^")
		visitValue(sb, f.Exponent)
	}
}

func visitOpFactor(sb *strings.Builder, of *OpFactor) {
}

func visitTerm(sb *strings.Builder, t *Term) {
	if t.Left != nil {
		visitFactor(sb, t.Left)
	}
	if t.Right != nil {
		for _, of := range t.Right {
			visitOpFactor(sb, of)
		}
	}
}

func visitOpTerm(sb *strings.Builder, o *OpTerm) {
	sb.WriteString(string(o.Operator))
	visitTerm(sb, o.Term)
}

func visitExpression(sb *strings.Builder, e *Expression) {
	if e.Left != nil {
		visitCmp(sb, e.Left)
	}
	if e.Right != nil {
		for _, o := range e.Right {
			visitOpCmp(sb, o)
		}
	}
}

func main() {
	parser, err := participle.Build[HiroAst]()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile("start.hi")
	if err != nil {
		panic(err)
	}
	hiro, err := parser.ParseString("", string(dat))
	if err != nil {
		panic(err)
	} else {
		spew.Dump(hiro)

		f, err := os.Create("target/my.go")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		var sb strings.Builder

		_, err = f.WriteString("package main\n\n")

		sb.WriteString("import (\n\t\"fmt\")\n\n")
		for _, fu := range hiro.Functions {
			sb.WriteString(fmt.Sprintf(`func %s(`, fu.Name))
			for i, arg := range fu.Args {
				sb.WriteString(arg.VarName)
				sb.WriteString(" ")
				sb.WriteString(arg.VarType)
				if i < len(fu.Args)-1 {
					sb.WriteString(", ")
				}
			}
			sb.WriteString(fmt.Sprintf(`) %s {`, fu.Return))
			sb.WriteString("\n")
			lastCommand := fu.Body[len(fu.Body)-1]
			if lastCommand.Expression != nil {
				sb.WriteString("return ")
			}
			for _, c := range fu.Body {
				visitCommand(&sb, c)
			}
			if fu.Name == "main" {
				sb.WriteString("fmt.Println(\"End\")\n")
			}
			sb.WriteString("}\n")
		}

		_, err = f.WriteString(sb.String())
	}
}

//func {:}() <- chan x {{
//	res := make(chan x)
//	go func() {{
//		defer close(res)
//		res <- x{{ 22 }}
//	}}()
//	return res
//}}
