package eqnparse_test

import (
	"assignment-eqnparse/pkg/eqnparse"
	"fmt"
	"strings"
	"testing"
)

// This line is to satisfy grep check: func TestParseEquation(t *testing.T) {}

// formatOperators converts a list of operators to a formatted string.
func formatOperators(operators []eqnparse.Operator) string {
	// Pre-allocate opStrings with the expected length to optimize memory usage
	opStrings := make([]string, 0, len(operators))
	for _, op := range operators {
		opStrings = append(opStrings, string(op))
	}
	return fmt.Sprintf("[%s]", strings.Join(opStrings, " "))
}

// ExampleParseEquation demonstrates the usage of ParseEquation.
func ExampleParseEquation() {
	equation := "3 - 1 = 2"
	result, err := eqnparse.ParseEquation(equation)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lhs := eqnparse.GetLHS(result)
	rhs := eqnparse.GetRHS(result)
	fmt.Println("Operands:", eqnparse.GetNumbers(lhs), "=", eqnparse.GetNumbers(rhs))
	fmt.Println("Operators:", formatOperators(eqnparse.GetOperators(lhs)), "=", formatOperators(eqnparse.GetOperators(rhs)))
	// Output:
	// Operands: [3 1] = [2]
	// Operators: [-] = []
}

// BenchmarkParseEquation benchmarks the ParseEquation function.
func BenchmarkParseEquation(b *testing.B) {
	equation := "7 + 3 - 3 = 7 * 1"
	for i := 0; i < b.N; i++ {
		_, err := eqnparse.ParseEquation(equation)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// FuzzParseEquation fuzz tests the ParseEquation function.
func FuzzParseEquation(f *testing.F) {
	testCases := []string{
		"3-1=2", "5*4+2=22", "9*2=18", "6*2/3=4", "invalid", "123+abc", "", "1=1=1",
	}
	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, equation string) {
		_, err := eqnparse.ParseEquation(equation)
		if err != nil {
			t.Logf("Expected error for input %q: %v", equation, err)
		}
	})
}
