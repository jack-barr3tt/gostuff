package set

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestFromSlice(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 2, 1})

	test.AssertEqual(t, s.Size(), 3)
	test.AssertEqual(t, s.Has(1), true)
	test.AssertEqual(t, s.Has(2), true)
	test.AssertEqual(t, s.Has(3), true)
	test.AssertEqual(t, s.Has(4), false)
}

func TestFromMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	s := FromMap(m)

	test.AssertEqual(t, s.Size(), 3)
	test.AssertEqual(t, s.Has("a"), true)
	test.AssertEqual(t, s.Has("b"), true)
	test.AssertEqual(t, s.Has("c"), true)
	test.AssertEqual(t, s.Has("d"), false)
}

func TestToSlice(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	slice := s.ToSlice()

	test.AssertSlicesEqual(t, slice, []int{1, 2, 3})
}

func TestAdd(t *testing.T) {
	s := FromSlice([]int{1, 2})

	test.AssertEqual(t, s.Size(), 2)
	test.AssertEqual(t, s.Has(3), false)

	s.Add(3)

	test.AssertEqual(t, s.Size(), 3)
	test.AssertEqual(t, s.Has(3), true)

	// Adding duplicate should not change size
	s.Add(3)
	test.AssertEqual(t, s.Size(), 3)
}

func TestRemove(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})

	test.AssertEqual(t, s.Size(), 3)
	test.AssertEqual(t, s.Has(2), true)

	s.Remove(2)

	test.AssertEqual(t, s.Size(), 2)
	test.AssertEqual(t, s.Has(2), false)

	// Removing non-existent item should not error
	s.Remove(4)
	test.AssertEqual(t, s.Size(), 2)
}

func TestHas(t *testing.T) {
	s := FromSlice([]string{"apple", "banana", "cherry"})

	test.AssertEqual(t, s.Has("apple"), true)
	test.AssertEqual(t, s.Has("banana"), true)
	test.AssertEqual(t, s.Has("cherry"), true)
	test.AssertEqual(t, s.Has("orange"), false)
	test.AssertEqual(t, s.Has(""), false)
}

func TestSize(t *testing.T) {
	s := FromSlice([]int{})
	test.AssertEqual(t, s.Size(), 0)

	s.Add(1)
	test.AssertEqual(t, s.Size(), 1)

	s.Add(2)
	s.Add(3)
	test.AssertEqual(t, s.Size(), 3)

	s.Remove(2)
	test.AssertEqual(t, s.Size(), 2)
}

func TestUnion(t *testing.T) {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{3, 4, 5})

	result := s1.Union(s2)

	test.AssertEqual(t, result.Size(), 5)
	test.AssertEqual(t, result.Has(1), true)
	test.AssertEqual(t, result.Has(2), true)
	test.AssertEqual(t, result.Has(3), true)
	test.AssertEqual(t, result.Has(4), true)
	test.AssertEqual(t, result.Has(5), true)

	// Original sets should not be modified
	test.AssertEqual(t, s1.Size(), 3)
	test.AssertEqual(t, s2.Size(), 3)
}

func TestIntersection(t *testing.T) {
	s1 := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{3, 4, 5, 6})

	result := s1.Intersection(s2)

	test.AssertEqual(t, result.Size(), 2)
	test.AssertEqual(t, result.Has(3), true)
	test.AssertEqual(t, result.Has(4), true)
	test.AssertEqual(t, result.Has(1), false)
	test.AssertEqual(t, result.Has(5), false)

	// Test with no common elements
	s3 := FromSlice([]int{1, 2})
	s4 := FromSlice([]int{3, 4})
	result2 := s3.Intersection(s4)
	test.AssertEqual(t, result2.Size(), 0)
}

func TestDifference(t *testing.T) {
	s1 := FromSlice([]int{1, 2, 3, 4})
	s2 := FromSlice([]int{3, 4, 5, 6})

	result := s1.Difference(s2)

	test.AssertEqual(t, result.Size(), 2)
	test.AssertEqual(t, result.Has(1), true)
	test.AssertEqual(t, result.Has(2), true)
	test.AssertEqual(t, result.Has(3), false)
	test.AssertEqual(t, result.Has(4), false)

	// Test with no overlap
	s3 := FromSlice([]int{1, 2})
	s4 := FromSlice([]int{3, 4})
	result2 := s3.Difference(s4)
	test.AssertEqual(t, result2.Size(), 2)
	test.AssertEqual(t, result2.Has(1), true)
	test.AssertEqual(t, result2.Has(2), true)
}

func TestIsSubset(t *testing.T) {
	s1 := FromSlice([]int{1, 2})
	s2 := FromSlice([]int{1, 2, 3, 4})

	test.AssertEqual(t, s1.IsSubset(s2), true)
	test.AssertEqual(t, s2.IsSubset(s1), false)

	// Test equal sets
	s3 := FromSlice([]int{1, 2, 3})
	s4 := FromSlice([]int{1, 2, 3})
	test.AssertEqual(t, s3.IsSubset(s4), true)

	// Test disjoint sets
	s5 := FromSlice([]int{1, 2})
	s6 := FromSlice([]int{3, 4})
	test.AssertEqual(t, s5.IsSubset(s6), false)

	// Test empty set is subset of any set
	s7 := FromSlice([]int{})
	test.AssertEqual(t, s7.IsSubset(s2), true)
}
