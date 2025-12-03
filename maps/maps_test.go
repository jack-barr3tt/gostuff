package maps

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestKeys(t *testing.T) {
	result1 := Keys(map[string]int{"a": 1, "b": 2})
	test.AssertSlicesEqual(t, result1, []string{"a", "b"})

	result2 := Keys(map[string]int{})
	test.AssertEqual(t, result2, []string{})

	result3 := Keys(map[int]string{1: "a", 2: "b"})
	test.AssertSlicesEqual(t, result3, []int{1, 2})

	result4 := Keys(map[int]string{})
	test.AssertEqual(t, result4, []int{})
}

func TestValues(t *testing.T) {
	result1 := Values(map[string]int{"a": 1, "b": 2})
	test.AssertSlicesEqual(t, result1, []int{1, 2})

	result2 := Values(map[string]int{})
	test.AssertEqual(t, result2, []int{})

	result3 := Values(map[int]string{1: "a", 2: "b"})
	test.AssertSlicesEqual(t, result3, []string{"a", "b"})

	result4 := Values(map[int]string{})
	test.AssertEqual(t, result4, []string{})
}

func TestClone(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	clone := Clone(m)
	test.AssertEqual(t, clone, m)

	m["c"] = 3
	test.AssertNotEqual(t, clone, m)
}

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	result := Map(func(k string, v int) int { return v * 2 }, m)
	test.AssertSlicesEqual(t, result, []int{2, 4})
}

func TestReduce(t *testing.T) {
	m := map[string]int{"aa": 1, "babd": 2, "cd": 3}
	result := Reduce(func(acc int, k string, v int) int {
		if len(k) > acc {
			return len(k)
		}
		return acc
	}, m, 0)
	test.AssertEqual(t, result, 4)
}
