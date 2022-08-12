package main

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

// BASIC example
// https://github.com/alecthomas/participle/blob/master/_examples/basic/ast.go

type HiroAst struct {
	Functions []*Function `@@*`
}

type Function struct {
	Name   string     `"fn" @Ident`
	Args   []*Arg     `"(" ( @@ ( "," @@ )* )? ")"`
	Return string     `("-" ">" @Ident)? ":" EOL`
	Body   []*Command `@@* "end" EOL`
}

type Command struct {
	Pos        lexer.Position
	Print      *Print `(@@ `
	Call       *Call
	Expression *Expression `| @@ ) EOL`
}

type Call struct {
	Pos lexer.Position

	Name string        `@Ident`
	Args []*Expression `"(" ( @@ ( "," @@ )* )? ")"`
}

type Print struct {
	Pos        lexer.Position
	Expression string `"print" @String`
}

type Arg struct {
	Pos     lexer.Position
	VarName string `@Ident ":"`
	VarType string `@Ident`
}

// Expressions

type Operator string

func (o *Operator) Capture(s []string) error {
	*o = Operator(strings.Join(s, ""))
	return nil
}

type Value struct {
	Pos lexer.Position

	Variable      *string     `  @Ident`
	String        *string     `| @String`
	Number        *int        `| @Int`
	Subexpression *Expression `| "(" @@ ")"`
}

type Factor struct {
	Pos lexer.Position

	Base     *Value `@@`
	Exponent *Value `( "^" @@ )?`
}

type OpFactor struct {
	Pos lexer.Position

	Operator Operator `@("*" | "/")`
	Factor   *Factor  `@@`
}

type Term struct {
	Pos lexer.Position

	Left  *Factor     `@@`
	Right []*OpFactor `@@*`
}

type OpTerm struct {
	Pos lexer.Position

	Operator Operator `@("+" | "-")`
	Term     *Term    `@@`
}

type Cmp struct {
	Pos lexer.Position

	Left  *Term     `@@`
	Right []*OpTerm `@@*`
}

type OpCmp struct {
	Pos lexer.Position

	Operator Operator `@("=" | "<" "=" | ">" "=" | "<" | ">" | "!" "=")`
	Cmp      *Cmp     `@@`
}

type Expression struct {
	Pos lexer.Position

	Left  *Cmp     `@@`
	Right []*OpCmp `@@*`
}
