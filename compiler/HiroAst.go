package compiler

import (
	"github.com/alecthomas/participle/v2/lexer"
)

// BASIC example
// https://github.com/alecthomas/participle/blob/master/_examples/basic/ast.go

type HiroAst struct {
	Functions []*Function `@@*`
}

type Function struct {
	Name   string     `"fn" @Ident`
	Args   []*Arg     `"(" ( @@ ( "," @@ )* )? ")"`
	Return string     `("-" ">" @Ident)? ":"`
	Body   []*Command `((?! 'end') @@)*`
	End    bool       `@"end"`
}

type Command struct {
	Pos lexer.Position

	Let        *Let        `( @@`
	Print      *Print      `| @@`
	Call       *Call       `| @@`
	Expression *Expression `| @@ )`
}

type Let struct {
	Pos lexer.Position

	Var string `"let" @Ident "="`
	//Args []string `"(" ( @String ( "," @String )* )? ")"`
	Expr *Expression `@@`
}

type Call struct {
	Pos lexer.Position

	Name string `@Ident`
	//Args []string `"(" ( @String ( "," @String )* )? ")"`
	Args []*Expression `"(" ( @@ ( "," @@ )* )? ")"`
}

type Print struct {
	Pos        lexer.Position
	Expression *Expression `"print" @@`
}

type Arg struct {
	Pos     lexer.Position
	VarName string `@Ident ":"`
	VarType string `@Ident`
}

// Expressions

type Expression struct {
	IsAsync         bool
	CheckedForAsync bool
	Equality        *Equality `@@`
}

type Equality struct {
	Comparison *Comparison `@@`
	Op         string      `[ @( "!" "=" | "=" "=" )`
	Next       *Equality   `  @@ ]`
}

type Comparison struct {
	Addition *Addition   `@@`
	Op       string      `[ @( ">" | ">" "=" | "<" | "<" "=" )`
	Next     *Comparison `  @@ ]`
}

type Addition struct {
	Multiplication *Multiplication `@@`
	Op             string          `[ @( "-" | "+" )`
	Next           *Addition       `  @@ ]`
}

type Multiplication struct {
	Unary *Unary          `@@`
	Op    string          `[ @( "/" | "*" )`
	Next  *Multiplication `  @@ ]`
}

type Unary struct {
	Op      string   `  ( @( "!" | "-" )`
	Unary   *Unary   `    @@ )`
	Primary *Primary `| @@`
}

type Primary struct {
	Call          *Call       `( @@ `
	Float         *float64    `| @Float`
	Int           *int        `| @Int`
	String        *string     `| @String`
	Bool          *bool       `| ( @"true" | "false" )`
	Nil           bool        `| @"nil"`
	Variable      *string     `| @Ident `
	SubExpression *Expression `| "(" @@ ")" )`
}
