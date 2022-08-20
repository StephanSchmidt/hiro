package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/alecthomas/participle/v2"
	"github.com/davecgh/go-spew/spew"
	"hiro/compiler"
	"os"
	"strings"
)

// Template for function to be called async
//
//func {:}() <- chan x {{
//	res := make(chan x)
//	go func() {{
//		defer close(res)
//		res <- x{{ 22 }}
//	}}()
//	return res
//}}

// Template for expression
//
//  a := 1 + 2
//
//	expr := make(chan int)
//	go func() {
//		defer close(expr)
//		expr <- 1 + 2
//	}()
//  ...
//  a := <- expr
//}

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

	argsParser := argparse.NewParser("hiro", "Hiro compiler")
	h := argsParser.String("s", "hiro", &argparse.Options{Required: true, Help: "Hiro source file"})
	t := argsParser.String("t", "target", &argparse.Options{Required: true, Help: "Target directory for go"})
	err := argsParser.Parse(os.Args)
	if err != nil {
		fmt.Print(argsParser.Usage(err))
	}

	parser, err := participle.Build[compiler.HiroAst](participle.UseLookahead(2))
	if err != nil {
		fmt.Println("Can't parse grammar.")
		panic(err)
	}

	dat, err := os.ReadFile(*h)
	if err != nil {
		panic(err)
	}
	hiro, err := parser.ParseString("", string(dat))
	if err != nil {
		spew.Dump(hiro)
		panic(err)
	} else {
		spew.Dump(hiro)

		var sb strings.Builder
		var symbols = compiler.NewSymbols()

		goGenerator := &compiler.GoGenerator{
			Sb:      &sb,
			Symbols: symbols,
		}

		f, err := os.Create(*t + "/my.go")
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		goGenerator.VisitAst(hiro)
		_, err = f.WriteString(sb.String())
	}
}
