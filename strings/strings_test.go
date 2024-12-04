package strings

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestGetNum(t *testing.T) {
	test.AssertCases(t, GetNum, []test.Case[string, int]{
		{Input: "1", Expected: 1},
		{Input: "abc123", Expected: 123},
		{Input: "abc-123", Expected: -123},
		{Input: "abc-123def456", Expected: -123},
		{Input: "", Expected: 0},
	})
}

func TestGetNums(t *testing.T) {
	test.AssertCases(t, GetNums, []test.Case[string, []int]{
		{Input: "1", Expected: []int{1}},
		{Input: "abc123", Expected: []int{123}},
		{Input: "abc-123", Expected: []int{-123}},
		{Input: "abc-123def456", Expected: []int{-123, 456}},
		{Input: "", Expected: []int{}},
	})
}