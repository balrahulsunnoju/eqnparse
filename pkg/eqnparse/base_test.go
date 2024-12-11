package eqnparse_test

import (
	"assignment-eqnparse/pkg/eqnparse"
	"reflect"
	"testing"
)

// STUDENTS: DO NOT MODIFY THIS FILE

type op = eqnparse.Operator

const (
	add = eqnparse.Addition
	sub = eqnparse.Subtraction
	mul = eqnparse.Multiplication
	div = eqnparse.Division
)

func buildEquation(lhsNums []int, lhsOps []op,
	rhsNums []int, rhsOps []op,
) eqnparse.Equation {
	return eqnparse.NewEquation(
		eqnparse.NewExpression(lhsNums, lhsOps),
		eqnparse.NewExpression(rhsNums, rhsOps),
	)
}

func TestParseEquation(t *testing.T) {
	tests := []struct {
		name    string
		eqnStr  string
		wantEqn eqnparse.Equation
	}{
		{
			"basic",
			" 7+ 2- 3= 6",
			buildEquation([]int{7, 2, 3}, []op{add, sub}, []int{6}, nil),
		},
		{
			"complex",
			"9* 3/4 =5 +2 ",
			buildEquation([]int{9, 3, 4}, []op{mul, div}, []int{5, 2}, []op{add}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEqn, err := eqnparse.ParseEquation(tt.eqnStr)
			if err != nil {
				t.Fatalf("Unexpected error %s", err)
			}

			if !reflect.DeepEqual(*gotEqn, tt.wantEqn) {
				t.Errorf("Expected %#v == %#v", gotEqn, tt.wantEqn)
			}
			if gotEqn.String() != tt.wantEqn.String() {
				t.Errorf("Expected %#v == %#v", gotEqn, tt.wantEqn)
			}
		})
	}
}
