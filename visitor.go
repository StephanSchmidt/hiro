package main

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
