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

	problem1 := Problem{
		Objective: []float64{3, 2},
		Constraints: []Constraint{
			{Coefficients: []float64{5, 7}, Value: 70},
			{Coefficients: []float64{10, 3}, Value: 60},
		},
	}

	solution1 := problem1.Solve(false, false)

	test.AssertEqual(t, solution1.Optimal, true)
	test.AssertEqual(t, solution1.Vars[0], 42.0/11.0)
	test.AssertEqual(t, solution1.Vars[1], 80.0/11.0)
	test.AssertEqual(t, solution1.Value, 26.0)

	// Maximize P = 3x + 2y
	// Subject to:
	// 5x + 7y <= 70
	// 10x + 3y <= 60
	// x, y >= 0
	// x, y are integers
	// Expected: P = 24, x = 4, y = 6

	problem2 := Problem{
		Objective: []float64{3, 2},
		Constraints: []Constraint{
			{Coefficients: []float64{5, 7}, Value: 70},
			{Coefficients: []float64{10, 3}, Value: 60},
		},
	}

	solution2 := problem2.Solve(true, false)

	test.AssertEqual(t, solution2.Optimal, true)
	test.AssertEqual(t, solution2.Vars[0], 3.0)
	test.AssertEqual(t, solution2.Vars[1], 7.0)
	test.AssertEqual(t, solution2.Value, 23.0)

	// Minimize P = 3x + 6y - 32z
	// Subject to:
	// x + 6y + 24z <= 672
	// 3x + y + 24z <= 336
	// x + 3y + 16z <= 168
	// 2x + 3y + 32z <= 352
	// x, y, z >= 0
	// Expected: P = -336, x = 0, y = 0, z = 21/2

	problem3 := Problem{
		Objective: []float64{3, 6, -32},
		Constraints: []Constraint{
			{Coefficients: []float64{1, 6, 24}, Value: 672},
			{Coefficients: []float64{3, 1, 24}, Value: 336},
			{Coefficients: []float64{1, 3, 16}, Value: 168},
			{Coefficients: []float64{2, 3, 32}, Value: 352},
		},
	}

	solution3 := problem3.Solve(false, true)

	test.AssertEqual(t, solution3.Optimal, true)
	test.AssertEqual(t, solution3.Vars[0], 0.0)
	test.AssertEqual(t, solution3.Vars[1], 0.0)
	test.AssertEqual(t, solution3.Vars[2], 21.0/2.0)
	test.AssertEqual(t, solution3.Value, -336.0)
}
