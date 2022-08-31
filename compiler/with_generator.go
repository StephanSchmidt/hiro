package compiler

import (
	"github.com/thoas/go-funk"
)

func IsEqualityOrComparison(e *Expression) bool {
	return len(e.Equality.Op) > 0 ||
		(e.Equality.Comparison != nil && len(e.Equality.Comparison.Op) > 0)
}

type WithType int32

const (
	Illegal WithType = iota
	Assertion
	UnitTest
	PropTest
)

type WithExpression struct {
	WithType          WithType
	Op                string
	LeftVars          []string
	LeftHasVars       bool
	LeftContainsVar   bool
	RightVars         []string
	RightHasVars      bool
	RightContainsVar  bool
	LeftContainsCall  bool
	RightContainsCall bool
}

func WithExpressionTypeFor(f *Function, we *WithExpression) WithType {
	noCalls := !(we.LeftContainsCall || we.RightContainsCall)
	noVars := !(we.LeftHasVars || we.RightHasVars)
	onlyRightCall := we.RightContainsCall && !we.LeftContainsCall
	onlyLeftCall := !we.RightContainsCall && we.LeftContainsCall
	// 1>2 // Illegal
	if noVars && noCalls {
		return Illegal
	}
	//  2 > add(2,1) // Illegal
	if noVars && onlyRightCall {
		return Illegal
	}
	// add(2,1) == 1 // UnitTest
	if noVars && onlyLeftCall {
		return UnitTest
	}
	// add(a,b) == a + b // PropTest
	if we.LeftHasVars && we.LeftContainsCall {
		return PropTest
	}
	// a > 1
	if !we.RightContainsVar && we.LeftContainsVar && noCalls {
		return Assertion
	}
	return Illegal
}

func ToWithExpression(rl *RightLeft, funcName string, vars []string) *WithExpression {
	leftVars := funk.Uniq(rl.LeftVars).([]string)
	rightVars := funk.Uniq(rl.RightVars).([]string)
	we := &WithExpression{
		LeftVars:          leftVars,
		RightVars:         rightVars,
		RightHasVars:      len(rightVars) > 0,
		LeftHasVars:       len(leftVars) > 0,
		LeftContainsCall:  rl.LeftCalls != nil && funk.Contains(rl.LeftCalls, funcName),
		RightContainsCall: rl.RightCalls != nil && funk.Contains(rl.RightCalls, funcName),
		LeftContainsVar:   len(funk.Intersect(leftVars, vars).([]string)) > 0,
		RightContainsVar:  len(funk.Intersect(rightVars, vars).([]string)) > 0,
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
