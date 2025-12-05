package ranges

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
	"github.com/jack-barr3tt/gostuff/types"
)

func TestCombineRanges(t *testing.T) {
	cases := []test.Case[[]types.Range, []types.Range]{
		{
			Input: []types.Range{
				{Start: 1, End: 3},
				{Start: 5, End: 7},
				{Start: 2, End: 6},
			},
			Expected: []types.Range{
				{Start: 1, End: 7},
			},
		},
		{
			Input: []types.Range{
				{Start: 10, End: 15},
				{Start: 20, End: 25},
				{Start: 12, End: 18},
				{Start: 30, End: 35},
			},
			Expected: []types.Range{
				{Start: 10, End: 18},
				{Start: 20, End: 25},
				{Start: 30, End: 35},
			},
		},
		{
			Input: []types.Range{
				{Start: 1, End: 2},
				{Start: 3, End: 4},
				{Start: 5, End: 6},
			},
			Expected: []types.Range{
				{Start: 1, End: 2},
				{Start: 3, End: 4},
				{Start: 5, End: 6},
			},
		},
	}

	test.AssertCases(t, CombineRanges, cases)
}
