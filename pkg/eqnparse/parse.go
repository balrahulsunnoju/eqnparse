package eqnparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseEquation processes the input equation string and converts it into a structured Equation or returns an error if parsing fails.
func ParseEquation(equation string) (*Equation, error) {
	cleanEquation := strings.ReplaceAll(equation, " ", "")

	// Separate the equation into left and right parts around the '=' symbol.
	parts := strings.Split(cleanEquation, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("equation must have exactly one '=' symbol: %w", errors.New("parsing error"))
	}

	lhs, err := parseExpression(parts[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing left side: %w", err)
	}
	rhs, err := parseExpression(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing right side: %w", err)
	}

	return &Equation{lhs: lhs, rhs: rhs}, nil
}

// parseExpression interprets one side of an equation string as an Expression struct.
func parseExpression(expression string) (Expression, error) {
	var numList []int
	var opList []Operator
	var currentNumber string

	for _, char := range expression {
		switch {
		case char >= '0' && char <= '9':
			currentNumber += string(char)
		case char == '+' || char == '-' || char == '*' || char == '/':
			if currentNumber == "" {
				return Expression{}, errors.New("operator without preceding number")
			}
			num, err := strconv.Atoi(currentNumber)
			if err != nil {
				return Expression{}, fmt.Errorf("invalid number '%s': %w", currentNumber, err)
			}
			numList = append(numList, num)
			opList = append(opList, Operator(char))
			currentNumber = ""
		default:
			return Expression{}, errors.New("invalid character in expression")
		}
	}

	if currentNumber != "" {
		num, err := strconv.Atoi(currentNumber)
		if err != nil {
			return Expression{}, fmt.Errorf("invalid trailing number '%s': %w", currentNumber, err)
		}
		numList = append(numList, num)
	}

	return Expression{numbers: numList, operators: opList}, nil
}

// GetLHS retrieves the left-hand side expression of an Equation.
func GetLHS(eqn *Equation) Expression {
	return eqn.lhs
}

// GetRHS retrieves the right-hand side expression of an Equation.
func GetRHS(eqn *Equation) Expression {
	return eqn.rhs
}

// GetNumbers returns the list of numbers in an Expression.
func GetNumbers(expr Expression) []int {
	return expr.numbers
}

// GetOperators returns the list of operators in an Expression.
func GetOperators(expr Expression) []Operator {
	return expr.operators
}
