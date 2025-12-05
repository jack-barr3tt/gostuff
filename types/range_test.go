package types

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

type ContainsInput struct {
	r   Range
	num int
}

func TestContains(t *testing.T) {
	test.AssertCases(t, func(input ContainsInput) bool {
		return input.r.Contains(input.num)
	}, []test.Case[ContainsInput, bool]{
		{Input: ContainsInput{Range{Start: 1, End: 10}, 5}, Expected: true},
		{Input: ContainsInput{Range{Start: 1, End: 10}, 1}, Expected: true},
		{Input: ContainsInput{Range{Start: 1, End: 10}, 10}, Expected: true},
		{Input: ContainsInput{Range{Start: 5, End: 10}, 3}, Expected: false},
		{Input: ContainsInput{Range{Start: 5, End: 10}, 15}, Expected: false},
	})
}

type ContainsRangeInput struct {
	r     Range
	other Range
}

func TestContainsRange(t *testing.T) {
	test.AssertCases(t, func(input ContainsRangeInput) bool {
		return input.r.ContainsRange(input.other)
	}, []test.Case[ContainsRangeInput, bool]{
		{Input: ContainsRangeInput{Range{Start: 1, End: 10}, Range{Start: 3, End: 7}}, Expected: true},
		{Input: ContainsRangeInput{Range{Start: 5, End: 10}, Range{Start: 5, End: 10}}, Expected: true},
		{Input: ContainsRangeInput{Range{Start: 5, End: 10}, Range{Start: 3, End: 7}}, Expected: false},
		{Input: ContainsRangeInput{Range{Start: 5, End: 10}, Range{Start: 7, End: 12}}, Expected: false},
		{Input: ContainsRangeInput{Range{Start: 5, End: 10}, Range{Start: 15, End: 20}}, Expected: false},
		{Input: ContainsRangeInput{Range{Start: 5, End: 7}, Range{Start: 1, End: 10}}, Expected: false},
	})
}

type SplitAfterInput struct {
	r   Range
	num int
}

type SplitAfterOutput struct {
	r1 Range
	r2 Range
}

func TestSplitAfter(t *testing.T) {
	test.AssertCases(t, func(input SplitAfterInput) SplitAfterOutput {
		r1, r2 := input.r.SplitAfter(input.num)
		return SplitAfterOutput{r1, r2}
	}, []test.Case[SplitAfterInput, SplitAfterOutput]{
		{Input: SplitAfterInput{Range{Start: 1, End: 10}, 5}, Expected: SplitAfterOutput{Range{Start: 1, End: 5}, Range{Start: 6, End: 10}}},
		{Input: SplitAfterInput{Range{Start: 1, End: 10}, 1}, Expected: SplitAfterOutput{Range{Start: 1, End: 1}, Range{Start: 2, End: 10}}},
		{Input: SplitAfterInput{Range{Start: 1, End: 10}, 10}, Expected: SplitAfterOutput{Range{Start: 1, End: 10}, Range{Start: 11, End: 10}}},
	})

	// Test panic case separately
	t.Run("number not in range", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic but didn't get one")
			}
		}()
		Range{Start: 5, End: 10}.SplitAfter(3)
	})
}

type SubtractRangeInput struct {
	a Range
	b Range
}

func TestSubtractRange(t *testing.T) {
	test.AssertCases(t, func(input SubtractRangeInput) []Range {
		return input.a.SubtractRange(input.b)
	}, []test.Case[SubtractRangeInput, []Range]{
		{Input: SubtractRangeInput{Range{Start: 5, End: 10}, Range{Start: 5, End: 10}}, Expected: []Range{}},
		{Input: SubtractRangeInput{Range{Start: 1, End: 10}, Range{Start: 4, End: 6}}, Expected: []Range{{Start: 1, End: 3}, {Start: 7, End: 10}}},
		{Input: SubtractRangeInput{Range{Start: 5, End: 7}, Range{Start: 1, End: 10}}, Expected: []Range{}},
		{Input: SubtractRangeInput{Range{Start: 1, End: 5}, Range{Start: 10, End: 15}}, Expected: []Range{{Start: 1, End: 5}}},
		{Input: SubtractRangeInput{Range{Start: 10, End: 15}, Range{Start: 1, End: 5}}, Expected: []Range{{Start: 10, End: 15}}},
		{Input: SubtractRangeInput{Range{Start: 1, End: 7}, Range{Start: 5, End: 10}}, Expected: []Range{{Start: 1, End: 4}}},
		{Input: SubtractRangeInput{Range{Start: 5, End: 15}, Range{Start: 1, End: 10}}, Expected: []Range{{Start: 11, End: 15}}},
		{Input: SubtractRangeInput{Range{Start: 5, End: 15}, Range{Start: 5, End: 8}}, Expected: []Range{{Start: 9, End: 15}}},
		{Input: SubtractRangeInput{Range{Start: 5, End: 15}, Range{Start: 12, End: 15}}, Expected: []Range{{Start: 5, End: 11}}},
	})
}

type AddRangeInput struct {
	a Range
	b Range
}

func TestAddRange(t *testing.T) {
	test.AssertCases(t, func(input AddRangeInput) *Range {
		return input.a.AddRange(input.b)
	}, []test.Case[AddRangeInput, *Range]{
		{Input: AddRangeInput{Range{Start: 1, End: 5}, Range{Start: 3, End: 8}}, Expected: &Range{Start: 1, End: 8}},
		{Input: AddRangeInput{Range{Start: 1, End: 5}, Range{Start: 6, End: 10}}, Expected: nil},
		{Input: AddRangeInput{Range{Start: 1, End: 5}, Range{Start: 7, End: 10}}, Expected: nil},
		{Input: AddRangeInput{Range{Start: 1, End: 10}, Range{Start: 3, End: 7}}, Expected: &Range{Start: 1, End: 10}},
		{Input: AddRangeInput{Range{Start: 5, End: 10}, Range{Start: 5, End: 10}}, Expected: &Range{Start: 5, End: 10}},
		{Input: AddRangeInput{Range{Start: 5, End: 7}, Range{Start: 1, End: 10}}, Expected: &Range{Start: 1, End: 10}},
		{Input: AddRangeInput{Range{Start: 1, End: 3}, Range{Start: 10, End: 15}}, Expected: nil},
		{Input: AddRangeInput{Range{Start: 6, End: 10}, Range{Start: 1, End: 5}}, Expected: nil},
	})
}

type OverlapsRangeInput struct {
	a Range
	b Range
}

func TestOverlapsRange(t *testing.T) {
	test.AssertCases(t, func(input OverlapsRangeInput) bool {
		return input.a.OverlapsRange(input.b)
	}, []test.Case[OverlapsRangeInput, bool]{
		{Input: OverlapsRangeInput{Range{Start: 1, End: 5}, Range{Start: 3, End: 8}}, Expected: true},
		{Input: OverlapsRangeInput{Range{Start: 1, End: 5}, Range{Start: 6, End: 10}}, Expected: false},
		{Input: OverlapsRangeInput{Range{Start: 1, End: 5}, Range{Start: 7, End: 10}}, Expected: false},
		{Input: OverlapsRangeInput{Range{Start: 1, End: 10}, Range{Start: 3, End: 7}}, Expected: true},
		{Input: OverlapsRangeInput{Range{Start: 5, End: 10}, Range{Start: 5, End: 10}}, Expected: true},
		{Input: OverlapsRangeInput{Range{Start: 5, End: 7}, Range{Start: 1, End: 10}}, Expected: true},
		{Input: OverlapsRangeInput{Range{Start: 1, End: 3}, Range{Start: 10, End: 15}}, Expected: false},
		{Input: OverlapsRangeInput{Range{Start: 6, End: 10}, Range{Start: 1, End: 5}}, Expected: false},
	})
}

func TestWidth(t *testing.T) {
	test.AssertCases(t, func(r Range) int {
		return r.Width()
	}, []test.Case[Range, int]{
		{Input: Range{Start: 1, End: 10}, Expected: 10},
		{Input: Range{Start: 5, End: 5}, Expected: 1},
		{Input: Range{Start: -3, End: 3}, Expected: 7},
	})
}
