package pkg

// Visitor Not really a visitor currently,
// we do not use accept as we know the subtypes
type Visitor interface {
	VisitAst(*HiroAst)
	visitFunction(*Function)
	visitCommand(*Command)
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
