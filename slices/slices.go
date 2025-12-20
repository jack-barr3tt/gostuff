package slices

import (
	"math"
	"strconv"
	"sync"

	"github.com/jack-barr3tt/gostuff/types"
)

// Map HOF
func Map[I any, O any](fn func(I) O, l []I) []O {
	out := make([]O, len(l))
	for i, v := range l {
		out[i] = fn(v)
	}
	return out
}

// Filter HOF
func Filter[I any](fn func(I) bool, l []I) []I {
	out := make([]I, 0)
	for _, v := range l {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce HOF
func Reduce[Curr any, Acc any](fn func(Curr, Acc) Acc, l []Curr, init Acc) Acc {
	out := init
	for _, v := range l {
		out = fn(v, out)
	}
	return out
}

// FindIndex HOF
func FindIndex[I any](fn func(I) bool, l []I) int {
	for i, v := range l {
		if fn(v) {
			return i
		}
	}
	return -1
}

// Function to convert a slice of strings to a slice of ints
func StrsToInts(l []string) []int {
	return Map(func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}, Filter(func(el string) bool {
		_, err := strconv.Atoi(el)
		return err == nil
	}, l))
}

// Zip together two slices
func Zip[A comparable, B comparable](a []A, b []B) []types.Pair[A, B] {
	length := int(math.Min(float64(len(a)), float64(len(b))))
	out := make([]types.Pair[A, B], length)
	for i := 0; i < length; i++ {
		out[i] = types.Pair[A, B]{First: a[i], Second: b[i]}
	}
	return out
}

// Unzip a slice of pairs into two slices
func Unzip[A comparable, B comparable](l []types.Pair[A, B]) ([]A, []B) {
	a := make([]A, len(l))
	b := make([]B, len(l))
	for i, v := range l {
		a[i] = v.First
		b[i] = v.Second
	}
	return a, b
}

// Flatten a slice of slices
func Flat[A any](l [][]A) []A {
	out := make([]A, 0)
	for _, v := range l {
		out = append(out, v...)
	}
	return out
}

// Apply a function to each element of a slice and flatten the result
func FlatMap[A any, B any](fn func(A) []B, l []A) []B {
	out := make([]B, 0)
	for _, v := range l {
		out = append(out, fn(v)...)
	}
	return out
}

// Pair every element of one slice with every element of another slice
func Combos[A, B comparable](a []A, b []B) []types.Pair[A, B] {
	return CombosMap(func(a A, b B) types.Pair[A, B] { return types.Pair[A, B]{First: a, Second: b} }, a, b)
}

// Pair every element of one slice with every element of another slice and apply a function to each pair
func CombosMap[A, B, C any](fn func(A, B) C, a []A, b []B) []C {
	out := make([]C, 0)
	for _, v := range a {
		for _, w := range b {
			out = append(out, fn(v, w))
		}
	}
	return out
}

// Generate all possible combos when elements of a slice are paied with themselves n times
func NCombos[A comparable](l []A, n int) [][]A {
	if n == 0 {
		return [][]A{{}}
	}
	if len(l) == 0 {
		return [][]A{}
	}
	out := make([][]A, 0)
	for _, v := range l {
		for _, w := range NCombos(l, n-1) {
			out = append(out, append([]A{v}, w...))
		}
	}
	return out
}

// Generate all possible combos when elements of a slice are paied with themselves n times without duplicates
func NCombosUnique[A comparable](l []A, n int) [][]A {
	if n == 0 {
		return [][]A{{}}
	}
	if len(l) == 0 {
		return [][]A{}
	}
	out := make([][]A, 0)
	for i, v := range l {
		for _, w := range NCombosUnique(l[i+1:], n-1) {
			out = append(out, append([]A{v}, w...))
		}
	}
	return out
}

// Checks if a subsequence at the end of the slice of any length repeats at least twice direct before it
func HasRepeatingSuffix[A comparable](l []A, minLength int) bool {
	if len(l) < minLength*2 {
		return false
	}
	for i := minLength; i < len(l)/2; i++ {
		if Equals(l[len(l)-i:], l[len(l)-2*i:len(l)-i]) {
			return true
		}
	}
	return false
}

// Check if any element of a slice satisfies a predicate
func Some[I any](fn func(I) bool, l []I) bool {
	for _, v := range l {
		if fn(v) {
			return true
		}
	}
	return false
}

// Check if a slice starts with another slice
func StartsWith[A comparable](l []A, prefix []A) bool {
	if len(l) < len(prefix) {
		return false
	}
	for i, v := range prefix {
		if v != l[i] {
			return false
		}
	}
	return true
}

// Check if two slices contain all the same elements
func Equals[A comparable](a []A, b []A) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Count the frequency of each element in a slice
func Frequency[A comparable](l []A) map[A]int {
	out := make(map[A]int)
	for _, v := range l {
		out[v]++
	}
	return out
}

// Remove an element from a slice by index
func RemoveAt[A any](l []A, x int) []A {
	if x < 0 || x >= len(l) {
		return l
	}
	result := make([]A, 0, len(l)-1)
	result = append(result, l[:x]...)
	return append(result, l[x+1:]...)
}

func ParallelMap[A any, B any](fn func(A) B, inputs []A, coreCount int) []B {
	var wg sync.WaitGroup
	results := make([][]B, coreCount)
	chunkSize := (len(inputs) + coreCount - 1) / coreCount

	// Launch workers
	for i := 0; i < coreCount; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(inputs) {
			end = len(inputs)
		}

		wg.Add(1)
		go func(index int, chunk []A) {
			defer wg.Done()
			workerResults := make([]B, len(chunk))
			for j, item := range chunk {
				workerResults[j] = fn(item)
			}
			results[index] = workerResults
		}(i, inputs[start:end])
	}

	wg.Wait()

	var combined []B
	for _, r := range results {
		combined = append(combined, r...)
	}

	return combined
}

func Unique[A comparable](l []A) []A {
	out := make([]A, 0)
	seen := make(map[A]bool)
	for _, v := range l {
		if _, ok := seen[v]; !ok {
			out = append(out, v)
			seen[v] = true
		}
	}
	return out
}

func Repeat[A any](value A, count int) []A {
	out := make([]A, count)
	for i := 0; i < count; i++ {
		out[i] = value
	}
	return out
}

func CountIf[A any](f func(v A) bool, l []A) int {
	c := 0
	for _, v := range l {
		if f(v) {
			c++
		}
	}
	return c
}

func Sum[T ~int | ~float64 | ~int64 | ~float32](l []T) T {
	var sum T
	for _, v := range l {
		sum += v
	}
	return sum
}

func Min[T ~int | ~float64 | ~int64 | ~float32](l []T) T {
	if len(l) == 0 {
		var zero T
		return zero
	}
	min := l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[T ~int | ~float64 | ~int64 | ~float32](l []T) T {
	if len(l) == 0 {
		var zero T
		return zero
	}
	max := l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}