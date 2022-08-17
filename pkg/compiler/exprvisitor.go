package pkg

type ExprVisitor interface {
	visitCall(*Call)
	visitPrimary(*Primary)
	visitUnary(*Unary)
	visitMultiplication(*Multiplication)
	visitAddition(*Addition)
	visitComparison(*Comparison)
	visitEquality(*Equality)
	visitExpression(*Expression)
}
