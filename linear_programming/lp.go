package lp

import (
	"math"
)

type Constraint struct {
	Coefficients []float64 // LHS
	Value        float64   // RHS
}

// Problem represents a linear programming problem in standard form
// Maximize: c^T * x
// Subject to: A * x <= b, x >= 0
type Problem struct {
	Objective   []float64
	Constraints []Constraint
}

type Solution struct {
	Optimal bool
	Value   float64
	Vars    []float64
}

func (p *Problem) Solve() Solution {
	numVars := len(p.Objective)
	tableau := buildTableau(p, numVars)

	optimal := simplex(tableau, numVars)
	if !optimal {
		return Solution{Optimal: false}
	}

	return extractSolution(tableau, numVars)
}

func buildTableau(p *Problem, numVars int) [][]float64 {
	numConstraints := len(p.Constraints)
	numCols := numVars + numConstraints + 1 // vars + slacks + RHS

	tableau := make([][]float64, numConstraints+1)
	for i, c := range p.Constraints {
		tableau[i] = make([]float64, numCols)
		copy(tableau[i], c.Coefficients)
		tableau[i][numVars+i] = 1       // slack variable
		tableau[i][numCols-1] = c.Value // RHS
	}

	tableau[numConstraints] = make([]float64, numCols)
	for j, coeff := range p.Objective {
		tableau[numConstraints][j] = -coeff
	}

	return tableau
}

func simplex(tableau [][]float64, numVars int) bool {
	numConstraints := len(tableau) - 1
	numCols := len(tableau[0])

	for {
		// Find most negative coefficient
		enteringCol := -1
		minCoeff := 0.0
		for i := 0; i < numCols-1; i++ {
			if tableau[numConstraints][i] < minCoeff {
				minCoeff = tableau[numConstraints][i]
				enteringCol = i
			}
		}

		// If no negative coefficients, we're done
		if enteringCol == -1 {
			return true
		}

		// Find leaving variable (minimum ratio test)
		leavingRow := -1
		minRatio := math.Inf(1)
		for i := 0; i < numConstraints; i++ {
			if tableau[i][enteringCol] > 1e-10 {
				ratio := tableau[i][numCols-1] / tableau[i][enteringCol]
				if ratio < minRatio && ratio >= 0 {
					minRatio = ratio
					leavingRow = i
				}
			}
		}

		// If no valid leaving variable, problem is unbounded
		if leavingRow == -1 {
			return false
		}

		// Perform pivot operation
		pivot(tableau, leavingRow, enteringCol)
	}
}

func pivot(tableau [][]float64, pivotRow, pivotCol int) {
	numCols := len(tableau[0])
	pivotVal := tableau[pivotRow][pivotCol]

	// Normalize pivot row
	for j := 0; j < numCols; j++ {
		tableau[pivotRow][j] /= pivotVal
	}

	// Eliminate pivot column in other rows
	for i := 0; i < len(tableau); i++ {
		if i != pivotRow {
			factor := tableau[i][pivotCol]
			for j := 0; j < numCols; j++ {
				tableau[i][j] -= factor * tableau[pivotRow][j]
			}
		}
	}
}

func extractSolution(tableau [][]float64, numVars int) Solution {
	numConstraints := len(tableau) - 1
	numCols := len(tableau[0])

	solution := Solution{
		Optimal: true,
		Value:   tableau[numConstraints][numCols-1],
		Vars:    make([]float64, numVars),
	}

	// Find values of original variables
	for j := 0; j < numVars; j++ {
		// Check if this is a basic variable
		basicRow := findBasicRow(tableau, j, numConstraints)
		if basicRow != -1 {
			solution.Vars[j] = tableau[basicRow][numCols-1]
		}
	}

	return solution
}

func findBasicRow(tableau [][]float64, col int, numConstraints int) int {
	basicRow := -1
	for i := 0; i < numConstraints; i++ {
		if math.Abs(tableau[i][col]-1) < 1e-10 {
			// Check if all other entries in column are 0
			allZero := true
			for k := 0; k < numConstraints; k++ {
				if k != i && math.Abs(tableau[k][col]) > 1e-10 {
					allZero = false
					break
				}
			}
			if allZero {
				basicRow = i
				break
			}
		}
	}
	return basicRow
}
