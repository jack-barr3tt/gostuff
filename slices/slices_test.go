package slices

import (
	"strconv"
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
	"github.com/jack-barr3tt/gostuff/types"
)

func TestMap(t *testing.T) {
	// test same input and output type
	test.AssertEqual(t, Map(func(x int) int { return x * 2 }, []int{1, 2, 3}), []int{2, 4, 6})

	// test different input and output type
	test.AssertEqual(t, Map(func(x string) int { return len(x) }, []string{"a", "ab", "abc"}), []int{1, 2, 3})

	// test empty input
	test.AssertEqual(t, Map(func(x string) string { return x + "test" }, []string{}), []string{})
}

func TestFilter(t *testing.T) {
	// test same input and output type
	test.AssertEqual(t, Filter(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4}), []int{2, 4})

	// test different input and output type
	test.AssertEqual(t, Filter(func(x string) bool { return len(x) > 1 }, []string{"a", "ab", "abc"}), []string{"ab", "abc"})

	// test empty input
	test.AssertEqual(t, Filter(func(x string) bool { return len(x) > 1 }, []string{}), []string{})
}

func TestReduce(t *testing.T) {
	// test same input and output type
	test.AssertEqual(t, Reduce(func(curr, acc int) int { return acc + (curr * 2) }, []int{1, 2, 3, 4}, 0), 20)

	// test different input and output type
	test.AssertEqual(t, Reduce(func(curr, acc string) string { return acc + curr }, []string{"a", "b", "c"}, ""), "abc")

	// test empty input
	test.AssertEqual(t, Reduce(func(curr, acc string) string { return acc + curr }, []string{}, ""), "")
}

func TestStrsToInts(t *testing.T) {
	test.AssertEqual(t, StrsToInts([]string{"1", "2", "3"}), []int{1, 2, 3})

	test.AssertEqual(t, StrsToInts([]string{"1", "a", "3"}), []int{1, 3})

	test.AssertEqual(t, StrsToInts([]string{}), []int{})
}

func TestZip(t *testing.T) {
	a := []int{1, 2, 3}
	b := []string{"a", "b", "c"}
	c := []float64{1.1, 2.2, 3.3}

	test.AssertEqual(t, Zip(a, b), []types.Pair[int, string]{{First: 1, Second: "a"}, {First: 2, Second: "b"}, {First: 3, Second: "c"}})

	test.AssertEqual(t, Zip(a, c), []types.Pair[int, float64]{{First: 1, Second: 1.1}, {First: 2, Second: 2.2}, {First: 3, Second: 3.3}})

	test.AssertEqual(t, Zip(b, c), []types.Pair[string, float64]{{First: "a", Second: 1.1}, {First: "b", Second: 2.2}, {First: "c", Second: 3.3}})

	test.AssertEqual(t, Zip(a, []string{}), []types.Pair[int, string]{})

	test.AssertEqual(t, Zip([]int{}, b), []types.Pair[int, string]{})

	test.AssertEqual(t, Zip([]int{}, []string{}), []types.Pair[int, string]{})
}

func TestFlat(t *testing.T) {
	test.AssertEqual(t, Flat([][]int{{1, 2}, {3, 4}}), []int{1, 2, 3, 4})

	test.AssertEqual(t, Flat([][]int{{1, 2}, {}, {3, 4}}), []int{1, 2, 3, 4})

	test.AssertEqual(t, Flat([][]int{}), []int{})

	test.AssertEqual(t, Flat([][]int{{}}), []int{})
}

func TestFlatMap(t *testing.T) {
	test.AssertEqual(t, FlatMap(func(x int) []int { return []int{x, x + 1} }, []int{1, 2, 3}), []int{1, 2, 2, 3, 3, 4})

	test.AssertEqual(t, FlatMap(func(x int) []int { return []int{} }, []int{1, 2, 3}), []int{})

	test.AssertEqual(t, FlatMap(func(x int) []int { return []int{x, x + 1} }, []int{}), []int{})
}

func TestCombos(t *testing.T) {
	// test equal lengths same type
	test.AssertEqual(t, Combos([]int{1, 2, 3}, []int{4, 5, 6}), []types.Pair[int, int]{
		{First: 1, Second: 4}, {First: 1, Second: 5}, {First: 1, Second: 6},
		{First: 2, Second: 4}, {First: 2, Second: 5}, {First: 2, Second: 6},
		{First: 3, Second: 4}, {First: 3, Second: 5}, {First: 3, Second: 6},
	})

	// test equal lengths different type
	test.AssertEqual(t, Combos([]int{1, 2, 3}, []string{"a", "b", "c"}), []types.Pair[int, string]{
		{First: 1, Second: "a"}, {First: 1, Second: "b"}, {First: 1, Second: "c"},
		{First: 2, Second: "a"}, {First: 2, Second: "b"}, {First: 2, Second: "c"},
		{First: 3, Second: "a"}, {First: 3, Second: "b"}, {First: 3, Second: "c"},
	})

	// test different lengths
	test.AssertEqual(t, Combos([]int{1, 2}, []string{"a", "b", "c"}), []types.Pair[int, string]{
		{First: 1, Second: "a"}, {First: 1, Second: "b"}, {First: 1, Second: "c"},
		{First: 2, Second: "a"}, {First: 2, Second: "b"}, {First: 2, Second: "c"},
	})

	// test one empty
	test.AssertEqual(t, Combos([]int{}, []string{"a", "b", "c"}), []types.Pair[int, string]{})
}

func TestCombosMap(t *testing.T) {
	// test equal lengths same type
	test.AssertEqual(t, CombosMap(
		func(a, b int) int { return a * b },
		[]int{1, 2, 3}, []int{4, 5, 6}),
		[]int{4, 5, 6, 8, 10, 12, 12, 15, 18},
	)

	// test equal lengths different type
	test.AssertEqual(t, CombosMap(
		func(a int, b string) string { return strconv.Itoa(a) + b },
		[]int{1, 2, 3}, []string{"a", "b", "c"}),
		[]string{"1a", "1b", "1c", "2a", "2b", "2c", "3a", "3b", "3c"},
	)

	// test different lengths
	test.AssertEqual(t, CombosMap(
		func(a int, b int) int { return a + b },
		[]int{5, 10}, []int{1, 2, 3}),
		[]int{6, 7, 8, 11, 12, 13},
	)
}

func TestSome(t *testing.T) {
	// test empty
	test.AssertEqual(t, Some(func(x int) bool { return x > 0 }, []int{}), false)

	// test true
	test.AssertEqual(t, Some(func(x int) bool { return x > 0 }, []int{1, 2, 3}), true)

	// test false
	test.AssertEqual(t, Some(func(x int) bool { return x > 0 }, []int{-1, -2, -3}), false)
}

func TestStartsWith(t *testing.T) {
	// test empty
	test.AssertEqual(t, StartsWith([]int{}, []int{}), true)

	// test empty prefix
	test.AssertEqual(t, StartsWith([]int{1, 2, 3}, []int{}), true)

	// test empty slice
	test.AssertEqual(t, StartsWith([]int{}, []int{1, 2, 3}), false)

	// test equal
	test.AssertEqual(t, StartsWith([]int{1, 2, 3}, []int{1, 2}), true)

	// test not equal
	test.AssertEqual(t, StartsWith([]int{1, 2, 3}, []int{2, 3}), false)
}

func TestEquals(t *testing.T) {
	// test empty
	test.AssertEqual(t, Equals([]int{}, []int{}), true)

	// test equal
	test.AssertEqual(t, Equals([]int{1, 2, 3}, []int{1, 2, 3}), true)

	// test not equal
	test.AssertEqual(t, Equals([]int{1, 2, 3}, []int{1, 2, 4}), false)

	// test different lengths
	test.AssertEqual(t, Equals([]int{1, 2, 3}, []int{1, 2}), false)
}