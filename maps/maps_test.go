package maps

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestKeys(t *testing.T) {
	test.AssertCases(t, Keys[string, int], []test.Case[map[string]int, []string]{
		{Input: map[string]int{"a": 1, "b": 2}, Expected: []string{"a", "b"}},
		{Input: map[string]int{}, Expected: []string{}},
	})

	test.AssertCases(t, Keys[int, string], []test.Case[map[int]string, []int]{
		{Input: map[int]string{1: "a", 2: "b"}, Expected: []int{1, 2}},
		{Input: map[int]string{}, Expected: []int{}},
	})
}

func TestValues(t *testing.T) {
	test.AssertCases(t, Values[string, int], []test.Case[map[string]int, []int]{
		{Input: map[string]int{"a": 1, "b": 2}, Expected: []int{1, 2}},
		{Input: map[string]int{}, Expected: []int{}},
	})

	test.AssertCases(t, Values[int, string], []test.Case[map[int]string, []string]{
		{Input: map[int]string{1: "a", 2: "b"}, Expected: []string{"a", "b"}},
		{Input: map[int]string{}, Expected: []string{}},
	})
}

func TestClone(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	clone := Clone(m)
	test.AssertEqual(t, clone, m)

	m["c"] = 3
	test.AssertNotEqual(t, clone, m)
}
