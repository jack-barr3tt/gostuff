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
	test.AssertEqual(t, solution2.Vars[0], 4.0)
	test.AssertEqual(t, solution2.Vars[1], 6.0)
	test.AssertEqual(t, solution2.Value, 24.0)

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

	// Minimize P = sum of all 10 variables
	// Subject to 10 equality constraints with 10 variables
	// From Advent of Code 2025, Day 10, Part 2, my actual problem input
	problem6 := Problem{
		Objective: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		Constraints: []Constraint{
			{Coefficients: []float64{1, 0, 1, 0, 0, 1, 1, 0, 1, 0}, Value: 60, Type: EQ},
			{Coefficients: []float64{1, 0, 1, 0, 0, 1, 1, 0, 1, 1}, Value: 69, Type: EQ},
			{Coefficients: []float64{1, 0, 1, 1, 0, 1, 0, 1, 1, 1}, Value: 66, Type: EQ},
			{Coefficients: []float64{0, 1, 0, 0, 0, 0, 0, 0, 1, 0}, Value: 12, Type: EQ},
			{Coefficients: []float64{1, 0, 0, 0, 1, 0, 1, 1, 1, 1}, Value: 74, Type: EQ},
			{Coefficients: []float64{1, 1, 0, 0, 0, 1, 0, 1, 0, 1}, Value: 41, Type: EQ},
			{Coefficients: []float64{1, 1, 1, 0, 1, 1, 0, 0, 1, 1}, Value: 76, Type: EQ},
			{Coefficients: []float64{0, 1, 1, 0, 1, 0, 1, 0, 1, 0}, Value: 59, Type: EQ},
			{Coefficients: []float64{1, 1, 1, 1, 0, 1, 1, 0, 1, 0}, Value: 77, Type: EQ},
			{Coefficients: []float64{1, 0, 1, 0, 1, 1, 1, 0, 0, 1}, Value: 81, Type: EQ},
		},
	}

	solution6 := problem6.Solve(true, true)

	test.AssertEqual(t, solution6.Optimal, true)
	test.AssertEqual(t, solution6.Value, 107.0)

	// Minimize P = sum of all 12 variables
	// Subject to 10 equality constraints with 12 variables
	problem7 := Problem{
		Objective: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		Constraints: []Constraint{
			{Coefficients: []float64{0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0}, Value: 57, Type: EQ},
			{Coefficients: []float64{0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0}, Value: 31, Type: EQ},
			{Coefficients: []float64{0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0}, Value: 44, Type: EQ},
			{Coefficients: []float64{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1}, Value: 54, Type: EQ},
			{Coefficients: []float64{0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1}, Value: 68, Type: EQ},
			{Coefficients: []float64{0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0}, Value: 54, Type: EQ},
			{Coefficients: []float64{1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0}, Value: 52, Type: EQ},
			{Coefficients: []float64{0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0}, Value: 48, Type: EQ},
			{Coefficients: []float64{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1}, Value: 62, Type: EQ},
			{Coefficients: []float64{0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0}, Value: 47, Type: EQ},
		},
	}

	solution7 := problem7.Solve(true, true)

	test.AssertEqual(t, solution7.Optimal, true)
}

func TestDeepCopy(t *testing.T) {
	original := Problem{
		Objective: []float64{1, 2, 3},
		Constraints: []Constraint{
			{Coefficients: []float64{1, 2, 3}, Value: 10, Type: LE},
			{Coefficients: []float64{4, 5, 6}, Value: 20, Type: GE},
			{Coefficients: []float64{7, 8, 9}, Value: 30, Type: EQ},
		},
	}

	copied := original.Clone()

	// Verify values are equal
	test.AssertEqual(t, len(copied.Objective), len(original.Objective))
	test.AssertEqual(t, len(copied.Constraints), len(original.Constraints))
	for i := range original.Objective {
		test.AssertEqual(t, copied.Objective[i], original.Objective[i])
	}
	for i := range original.Constraints {
		test.AssertEqual(t, copied.Constraints[i].Value, original.Constraints[i].Value)
		test.AssertEqual(t, copied.Constraints[i].Type, original.Constraints[i].Type)
		test.AssertEqual(t, len(copied.Constraints[i].Coefficients), len(original.Constraints[i].Coefficients))
		for j := range original.Constraints[i].Coefficients {
			test.AssertEqual(t, copied.Constraints[i].Coefficients[j], original.Constraints[i].Coefficients[j])
		}
	}

	// Modify copied and verify original is unchanged
	copied.Objective[0] = 999
	copied.Constraints[0].Value = 999
	copied.Constraints[0].Coefficients[0] = 999

	test.AssertEqual(t, original.Objective[0], 1.0)
	test.AssertEqual(t, original.Constraints[0].Value, 10.0)
	test.AssertEqual(t, original.Constraints[0].Coefficients[0], 1.0)
}
