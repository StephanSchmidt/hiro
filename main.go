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
		// sb.WriteString(c.Print.Expression)
		sb.WriteString(")")
		sb.WriteString("\n")
	}
	if c.Expression != nil {
		visitExpression(sb, c.Expression)
		sb.WriteString("\n")
	}
	if c.Call != nil {
		visitCall(sb, c.Call)
		sb.WriteString("\n")
	}
	if c.Let != nil {
		visitLet(sb, c.Let)
		sb.WriteString("\n")
	}
}

func visitLet(sb *strings.Builder, l *Let) {
	sb.WriteString("var ")
	sb.WriteString(l.Var)
	sb.WriteString(" = ")
	visitExpression(sb, l.Expression)
}

func visitCall(sb *strings.Builder, c *Call) {
	sb.WriteString(c.Name)
	sb.WriteString("(")
	for i, par := range c.Args {
		//sb.WriteString(par)
		visitExpression(sb, par)
		if i < len(c.Args)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString(")")
}

func visitPrimary(sb *strings.Builder, p *Primary) {
	if p.String != nil {
		sb.WriteString(*p.String)
	}
	if p.Float != nil {
		sb.WriteString(strconv.FormatFloat(3.1415, 'E', -1, 64))
	}
	if p.Int != nil {
		sb.WriteString(strconv.Itoa(*p.Int))
	}
	if p.Variable != nil {
		sb.WriteString(*p.Variable)
	}
	if p.SubExpression != nil {
		sb.WriteString("(")
		visitExpression(sb, p.SubExpression)
		sb.WriteString(")")
	}
}

func visitUnary(sb *strings.Builder, u *Unary) {
	sb.WriteString(u.Op)
	if u.Unary != nil {
		visitUnary(sb, u.Unary)
	}
	visitPrimary(sb, u.Primary)
}

func visitMultiplication(sb *strings.Builder, m *Multiplication) {
	if m.Unary != nil {
		visitUnary(sb, m.Unary)
	}
	sb.WriteString(m.Op)
	if m.Next != nil {
		visitMultiplication(sb, m.Next)
	}
}

func visitAddition(sb *strings.Builder, a *Addition) {
	visitMultiplication(sb, a.Multiplication)
	sb.WriteString(a.Op)
	if a.Next != nil {
		visitAddition(sb, a.Next)
	}
}

func visitComparison(sb *strings.Builder, c *Comparison) {
	visitAddition(sb, c.Addition)
	sb.WriteString(c.Op)
	if c.Next != nil {
		visitComparison(sb, c.Next)
	}

}

func visitEquality(sb *strings.Builder, e *Equality) {
	visitComparison(sb, e.Comparison)
	sb.WriteString(e.Op)
	if e.Next != nil {
		visitEquality(sb, e.Next)
	}
}

func visitExpression(sb *strings.Builder, e *Expression) {
	if e.Equality != nil {
		visitEquality(sb, e.Equality)
	}
}

func main() {
	//basicLexer := lexer.MustSimple([]lexer.SimpleRule{
	//	{"Comment", `(?i)rem[^\n]*`},
	//	{"Int", `[-+]?\d+`},
	//	{"String", `"(\\"|[^"])+"`},
	//	{"Number", `[-+]?(\d*\.)?\d+`},
	//	{"Ident", `[a-zA-Z_]\w*`},
	//	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	//	{"EOL", `[\n\r]+`},
	//	{"whitespace", `[ \t]+`},
	//})

	var parser, err = participle.Build[HiroAst](participle.UseLookahead(2))
	if err != nil {
		fmt.Println("Can't parse grammar.")
		panic(err)
	}

	dat, err := os.ReadFile("start.hi")
	if err != nil {
		panic(err)
	}
	hiro, err := parser.ParseString("", string(dat))
	if err != nil {
		spew.Dump(hiro)
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
			for index, c := range fu.Body {
				if index == len(fu.Body)-1 && c.Expression != nil {
					sb.WriteString("return ")

				}
				visitCommand(&sb, c)
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
