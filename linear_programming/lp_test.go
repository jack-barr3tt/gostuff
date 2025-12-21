package lp

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestSimplex(t *testing.T) {
	// Maximize P = 3x + 2y
	// Subject to:
	// 5x + 7y <= 70
	// 10x + 3y <= 60
	// x, y >= 0
	// Expected: P = 26, x = 42/11, y = 80/11
	// From: Pearson Edexcel A-Level Decision Mathematics 1, Chapter 7 Example 6

	problem := Problem{
		Objective: []float64{3, 2},
		Constraints: []Constraint{
			{Coefficients: []float64{5, 7}, Value: 70},
			{Coefficients: []float64{10, 3}, Value: 60},
		},
	}

	expectedX := 42.0 / 11.0
	expectedY := 80.0 / 11.0
	expectedP := 26.0

	solution := problem.Solve()

	test.AssertEqual(t, solution.Optimal, true)
	test.AssertEqual(t, solution.Vars[0], expectedX)
	test.AssertEqual(t, solution.Vars[1], expectedY)
	test.AssertEqual(t, solution.Value, expectedP)
}
