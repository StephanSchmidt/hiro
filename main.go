package main

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/davecgh/go-spew/spew"
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

		var sb strings.Builder
		goGenerator := &GoGenerator{
			sb: &sb,
		}

		f, err := os.Create("target/my.go")
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		goGenerator.visitAst(hiro)
		_, err = f.WriteString(sb.String())
	}
}
