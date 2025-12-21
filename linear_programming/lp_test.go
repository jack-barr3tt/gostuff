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
			{Coefficients: []float64{5, 7}, Value: 70, Type: LE},
			{Coefficients: []float64{10, 3}, Value: 60, Type: LE},
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
			{Coefficients: []float64{5, 7}, Value: 70, Type: LE},
			{Coefficients: []float64{10, 3}, Value: 60, Type: LE},
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
			{Coefficients: []float64{1, 6, 24}, Value: 672, Type: LE},
			{Coefficients: []float64{3, 1, 24}, Value: 336, Type: LE},
			{Coefficients: []float64{1, 3, 16}, Value: 168, Type: LE},
			{Coefficients: []float64{2, 3, 32}, Value: 352, Type: LE},
		},
	}

	solution3 := problem3.Solve(false, true)

	test.AssertEqual(t, solution3.Optimal, true)
	test.AssertEqual(t, solution3.Vars[0], 0.0)
	test.AssertEqual(t, solution3.Vars[1], 0.0)
	test.AssertEqual(t, solution3.Vars[2], 21.0/2.0)
	test.AssertEqual(t, solution3.Value, -336.0)

	// Minimize P = x - y + z
	// Subject to:
	// 2x + y + z <= 20
	// x - 2y - z <= 7
	// x >= 4
	// x, y, z >= 0
	// Expected: P = -8, x = 4, y = 12, z = 0

	problem4 := Problem{
		Objective: []float64{1, -1, 1},
		Constraints: []Constraint{
			{Coefficients: []float64{2, 1, 1}, Value: 20, Type: LE},
			{Coefficients: []float64{1, -2, -1}, Value: 7, Type: LE},
			{Coefficients: []float64{1, 0, 0}, Value: 4, Type: GE},
		},
	}

	solution4 := problem4.Solve(false, true)

	test.AssertEqual(t, solution4.Optimal, true)
	test.AssertEqual(t, solution4.Vars[0], 4.0)
	test.AssertEqual(t, solution4.Vars[1], 12.0)
	test.AssertEqual(t, solution4.Vars[2], 0.0)
	test.AssertEqual(t, solution4.Value, -8.0)

	// Minimize P = a + b + c + d + e + f
	// Subject to:
	// e + f = 3
	// b + f = 5
	// c + d + e = 4
	// a + b + d = 7
	// a, b, c, d, e, f >= 0
	// Expected: P = 10

	problem5 := Problem{
		Objective: []float64{1, 1, 1, 1, 1, 1},
		Constraints: []Constraint{
			{Coefficients: []float64{0, 0, 0, 0, 1, 1}, Value: 3, Type: EQ},
			{Coefficients: []float64{0, 1, 0, 0, 0, 1}, Value: 5, Type: EQ},
			{Coefficients: []float64{0, 0, 1, 1, 1, 0}, Value: 4, Type: EQ},
			{Coefficients: []float64{1, 1, 0, 1, 0, 0}, Value: 7, Type: EQ},
		},
	}

	solution5 := problem5.Solve(false, true)

	test.AssertEqual(t, solution5.Optimal, true)
	test.AssertEqual(t, solution5.Value, 10.0)
}
