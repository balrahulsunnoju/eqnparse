package eqnparse

// STUDENTS: DO NOT MODIFY THIS FILE

// Operator represents mathematical binary operators
type Operator rune

// These are the valid mathemematical operators
const (
	Addition       Operator = '+'
	Subtraction    Operator = '-'
	Multiplication Operator = '*'
	Division       Operator = '/'
)

// ValidOperators is the list of valid operators for an expression
var ValidOperators = []Operator{Addition, Subtraction, Multiplication, Division}

// Expression represents a mathematical expression
type Expression struct {
	numbers   []int
	operators []Operator
}

// NewExpression builds an expression from numbers and operators
func NewExpression(numbers []int, operators []Operator) Expression {
	if len(numbers) != len(operators)+1 {
		panic("Inconsistent length of numbers and operators for an expression.  One more number than operators is expected.")
	}
	return Expression{
		numbers:   numbers,
		operators: operators,
	}
}

// Equation represents a mathematical equation
type Equation struct {
	lhs, rhs Expression
}

// NewEquation builds an equation from the left and right hand sides
func NewEquation(lhs, rhs Expression) Equation {
	return Equation{
		lhs: lhs,
		rhs: rhs,
	}
}
