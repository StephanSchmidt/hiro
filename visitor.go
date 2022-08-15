package main

// Visitor Not really a visitor currently,
// we do not use accept as we know the subtypes
type Visitor interface {
	visitAst(*HiroAst)
	visitFunction(*Function)
	visitCommand(*Command)
	visitExpr(*Expr)
	visitLet(*Let)
	visitCall(*Call)
	visitPrimary(*Primary)
	visitUnary(*Unary)
	visitMultiplication(*Multiplication)
	visitAddition(*Addition)
	visitComparison(*Comparison)
	visitEquality(*Equality)
	visitExpression(*Expression)
}
