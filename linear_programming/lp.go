package lp

import (
	"math"
	"sort"

	slicestuff "github.com/jack-barr3tt/gostuff/slices"
)

type ConstraintType string

var M = 1e6 // Large penalty value

const (
	LE ConstraintType = "<="
	GE ConstraintType = ">="
	EQ ConstraintType = "="
)

type Constraint struct {
	Coefficients []float64 // LHS
	Value        float64   // RHS
	Type         ConstraintType
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

func (p *Problem) Solve(requireIntegers bool, minimize bool) Solution {
	originalObjective := make([]float64, len(p.Objective))
	copy(originalObjective, p.Objective)

	// For minimization, negate the objective function
	if minimize {
		p.Objective = slicestuff.Map(func(v float64) float64 { return -v }, p.Objective)
	}

	numVars := len(p.Objective)
	tableau := buildTableau(p, numVars)

	optimal := simplex(tableau, numVars)
	if !optimal {
		return Solution{Optimal: false}
	}

	solution := extractSolution(tableau, numVars)

	if requireIntegers {
		solution = p.findIntegerSolution(solution)
	}

	// Restore original objective and negate solution value if minimizing
	p.Objective = originalObjective
	if minimize {
		solution.Value = -solution.Value
	}

	return solution
}

func buildTableau(p *Problem, numVars int) [][]float64 {
	numConstraints := len(p.Constraints)
	numArtificial := slicestuff.CountIf(func(c Constraint) bool { return c.Type == GE || c.Type == EQ }, p.Constraints)
	numSlackSurplus := slicestuff.CountIf(func(c Constraint) bool { return c.Type != EQ }, p.Constraints)
	numCols := numVars + numSlackSurplus + numArtificial + 1

	tableau := make([][]float64, numConstraints+1)

	slackIdx := 0
	artificialIdx := 0
	for i, c := range p.Constraints {
		tableau[i] = make([]float64, numCols)
		copy(tableau[i], c.Coefficients)

		if c.Type == GE {
			// For >= subtract surplus variable, add artificial variable
			tableau[i][numVars+slackIdx] = -1                     // surplus variable
			tableau[i][numVars+numSlackSurplus+artificialIdx] = 1 // artificial variable
			slackIdx++
			artificialIdx++
		} else if c.Type == EQ {
			// For = only add artificial variable (no slack/surplus)
			tableau[i][numVars+numSlackSurplus+artificialIdx] = 1 // artificial variable
			artificialIdx++
		} else {
			// For <= add slack variable
			tableau[i][numVars+slackIdx] = 1
			slackIdx++
		}

		tableau[i][numCols-1] = c.Value
	}

	tableau[numConstraints] = make([]float64, numCols)
	for j, coeff := range p.Objective {
		tableau[numConstraints][j] = -coeff
	}

	artificialIdx = 0
	for i, c := range p.Constraints {
		if c.Type == GE || c.Type == EQ {
			artCol := numVars + numSlackSurplus + artificialIdx
			tableau[numConstraints][artCol] = M

			for j := 0; j < numCols; j++ {
				tableau[numConstraints][j] -= M * tableau[i][j]
			}

			artificialIdx++
		}
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

func applyRounding(vals []float64, roundUp []bool) []float64 {
	rounded := make([]float64, len(vals))
	for i, v := range vals {
		if roundUp[i] {
			rounded[i] = math.Ceil(v)
		} else {
			rounded[i] = math.Floor(v)
		}
	}
	return rounded
}

func calculateValue(p *Problem, vals []float64) float64 {
	value := 0.0
	for i, v := range vals {
		value += float64(v) * p.Objective[i]
	}
	return value
}

func (p *Problem) isFeasible(vals []float64) bool {
	for _, c := range p.Constraints {
		sum := 0.0
		for i, coeff := range c.Coefficients {
			sum += coeff * vals[i]
		}

		switch c.Type {
		case LE:
			if sum > c.Value+1e-10 {
				return false
			}
		case GE:
			if sum < c.Value-1e-10 {
				return false
			}
		case EQ:
			if math.Abs(sum-c.Value) > 1e-10 {
				return false
			}
		}
	}
	return true
}

func (p *Problem) findIntegerSolution(initial Solution) Solution {
	numVars := len(p.Objective)

	roundingCombos := slicestuff.NCombos([]bool{true, false}, numVars)
	sort.Slice(roundingCombos, func(a, b int) bool {
		aCoeffs, bCoeffs := applyRounding(initial.Vars, roundingCombos[a]), applyRounding(initial.Vars, roundingCombos[b])

		return calculateValue(p, aCoeffs) > calculateValue(p, bCoeffs)
	})
	roundingCombos = slicestuff.Filter(func(combo []bool) bool {
		newVars := applyRounding(initial.Vars, combo)
		return p.isFeasible(newVars)
	}, roundingCombos)

	newVars := applyRounding(initial.Vars, roundingCombos[0])

	solution := Solution{
		Optimal: true,
		Vars:    newVars,
		Value:   calculateValue(p, newVars),
	}

	return solution
}
