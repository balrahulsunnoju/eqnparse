package eqnparse

import (
	"strconv"
	"strings"
)

// This String converts an Expression to a formatted string.
func (expr Expression) String() string {
	var builder strings.Builder
	totalLength := len(expr.numbers) + len(expr.operators)
	builder.Grow(totalLength * 2)

	// Alternately add numbers and operators to form the expression.
	for i := 0; i < totalLength; i++ {
		if i%2 == 0 {
			if _, err := builder.WriteString(strconv.Itoa(expr.numbers[i/2])); err != nil {
				return ""
			}
		} else {
			if _, err := builder.WriteRune(rune(expr.operators[(i-1)/2])); err != nil {
				return ""
			}
		}
	}
	return builder.String()
}

// This is a String method for Equation returns a string in the form "LHS = RHS".
func (eqn Equation) String() string {
	var builder strings.Builder
	builder.Grow(len(eqn.lhs.String()) + len(eqn.rhs.String()) + 1)

	if _, err := builder.WriteString(eqn.lhs.String()); err != nil {
		return ""
	}
	if _, err := builder.WriteString("="); err != nil {
		return ""
	}
	if _, err := builder.WriteString(eqn.rhs.String()); err != nil {
		return ""
	}

	return builder.String()
}
