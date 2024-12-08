package slices

import (
	"strconv"
	"testing"
	"time"

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

func TestFindIndex(t *testing.T) {
	// test empty
	test.AssertEqual(t, FindIndex(func(x int) bool { return x > 0 }, []int{}), -1)

	// test found
	test.AssertEqual(t, FindIndex(func(x int) bool { return x > 0 }, []int{-1, 0, 1, 2}), 2)

	// test not found
	test.AssertEqual(t, FindIndex(func(x int) bool { return x > 0 }, []int{-1, 0, -2}), -1)
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

func TestNCombos(t *testing.T) {
	// test n = 0
	test.AssertEqual(t, NCombos([]int{1, 2, 3}, 0), [][]int{{}})

	// test empty
	test.AssertEqual(t, NCombos([]int{}, 1), [][]int{})

	// test n = 1
	test.AssertEqual(t, NCombos([]int{1, 2, 3}, 1), [][]int{{1}, {2}, {3}})

	// test n = 2
	test.AssertEqual(t, NCombos([]int{1, 2, 3}, 2), [][]int{
		{1, 1}, {1, 2}, {1, 3},
		{2, 1}, {2, 2}, {2, 3},
		{3, 1}, {3, 2}, {3, 3},
	})

	// test n = 3
	test.AssertEqual(t, NCombos([]int{1, 2, 3}, 3), [][]int{
		{1, 1, 1}, {1, 1, 2}, {1, 1, 3},
		{1, 2, 1}, {1, 2, 2}, {1, 2, 3},
		{1, 3, 1}, {1, 3, 2}, {1, 3, 3},
		{2, 1, 1}, {2, 1, 2}, {2, 1, 3},
		{2, 2, 1}, {2, 2, 2}, {2, 2, 3},
		{2, 3, 1}, {2, 3, 2}, {2, 3, 3},
		{3, 1, 1}, {3, 1, 2}, {3, 1, 3},
		{3, 2, 1}, {3, 2, 2}, {3, 2, 3},
		{3, 3, 1}, {3, 3, 2}, {3, 3, 3},
	})
}

func TestNCombosUnique(t *testing.T) {
	// test n = 0
	test.AssertEqual(t, NCombosUnique([]int{1, 2, 3}, 0), [][]int{{}})

	// test empty
	test.AssertEqual(t, NCombosUnique([]int{}, 1), [][]int{})

	// test n = 1
	test.AssertEqual(t, NCombosUnique([]int{1, 2, 3}, 1), [][]int{{1}, {2}, {3}})

	// test n = 2
	test.AssertEqual(t, NCombosUnique([]int{1, 2, 3}, 2), [][]int{
		{1, 2}, {1, 3},
		{2, 3},
	})

	// test n = 3
	test.AssertEqual(t, NCombosUnique([]int{1, 2, 3}, 3), [][]int{
		{1, 2, 3},
	})
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

func TestHasRepeatingSuffix(t *testing.T) {
	// test empty
	test.AssertEqual(t, HasRepeatingSuffix([]int{}, 1), false)

	// test no repeating suffix
	test.AssertEqual(t, HasRepeatingSuffix([]int{1, 2, 3, 4, 5}, 1), false)

	// test repeating suffix
	test.AssertEqual(t, HasRepeatingSuffix([]int{1, 2, 3, 4, 5, 1, 2, 1, 2}, 1), true)

	// test no repeating suffix
	test.AssertEqual(t, HasRepeatingSuffix([]int{1, 2, 3, 4, 5, 1, 2, 1, 2, 2}, 2), false)
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

func TestFrequency(t *testing.T) {
	// test empty
	test.AssertEqual(t, Frequency([]int{}), map[int]int{})

	// test one element
	test.AssertEqual(t, Frequency([]int{1}), map[int]int{1: 1})

	// test multiple elements
	test.AssertEqual(t, Frequency([]int{1, 1, 2, 3, 3, 3}), map[int]int{1: 2, 2: 1, 3: 3})
}

func TestRemoveAt(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	// check element removed
	test.AssertEqual(t, RemoveAt(a, 2), []int{1, 2, 4, 5})
	// check that the original slice is not modified
	test.AssertEqual(t, a, []int{1, 2, 3, 4, 5})

	// test empty
	test.AssertEqual(t, RemoveAt([]int{}, 0), []int{})
}

func TestParallelMap(t *testing.T) {
	// test same input and output type
	test.AssertEqual(t, ParallelMap(func(x int) int { return x * 2 }, []int{1, 2, 3}, 4), []int{2, 4, 6})

	// test different input and output type
	test.AssertEqual(t, ParallelMap(func(x string) int { return len(x) }, []string{"a", "ab", "abc"}, 4), []int{1, 2, 3})

	testFunc := func(x int) int {
		time.Sleep(time.Second)
		return x * 2
	}

	// test with 1 worker
	start := time.Now()
	ParallelMap(testFunc, []int{1, 2, 3, 4, 5}, 1)
	test.AssertEqual(t, time.Since(start) > 5*time.Second, true)

	// test with 5 workers
	start = time.Now()
	ParallelMap(testFunc, []int{1, 2, 3, 4, 5}, 5)
	test.AssertEqual(t, time.Since(start) < 2*time.Second, true)
}

func TestUnique(t *testing.T) {
	// test empty
	test.AssertEqual(t, Unique([]int{}), []int{})

	// test no duplicates
	test.AssertEqual(t, Unique([]int{1, 2, 3}), []int{1, 2, 3})

	// test duplicates
	test.AssertEqual(t, Unique([]int{1, 2, 1, 3, 2, 3}), []int{1, 2, 3})
}
