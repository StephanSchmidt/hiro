package compiler

import "github.com/thoas/go-funk"

func IsEqualityOrComparison(e *Expression) bool {
	return len(e.Equality.Op) > 0 ||
		(e.Equality.Comparison != nil && len(e.Equality.Comparison.Op) > 0)
}

type WithExpression struct {
	Op                string
	LeftVars          []string
	RightVars         []string
	LeftContainsCall  bool
	RightContainsCall bool
}

func ToWithExpression(rl *RightLeft, funcName string) *WithExpression {
	we := &WithExpression{
		LeftVars:          funk.Uniq(rl.LeftVars).([]string),
		RightVars:         funk.Uniq(rl.RightVars).([]string),
		LeftContainsCall:  rl.LeftCalls != nil && funk.Contains(rl.LeftCalls, funcName),
		RightContainsCall: rl.RightCalls != nil && funk.Contains(rl.RightCalls, funcName),
		Op:                rl.Op,
	}
	return we
}

type RightLeft struct {
	Op         string
	LeftVars   []string
	LeftCalls  []string
	RightVars  []string
	RightCalls []string
}

func RightLeftVars(e *Expression) *RightLeft {
	rightLeft := &RightLeft{}
	left := &ExpressionAnalyzer{}
	right := &ExpressionAnalyzer{}
	if len(e.Equality.Op) > 0 {
		rightLeft.Op = e.Equality.Op
		left.visitComparison(e.Equality.Comparison)
		right.visitEquality(e.Equality.Next)
	} else {
		rightLeft.Op = e.Equality.Comparison.Op
		left.visitAddition(e.Equality.Comparison.Addition)
		right.visitComparison(e.Equality.Comparison.Next)
	}
	rightLeft.LeftVars = left.Vars
	rightLeft.LeftCalls = left.Calls
	rightLeft.RightVars = right.Vars
	rightLeft.RightCalls = right.Calls
	return rightLeft
}