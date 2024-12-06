package slices

import (
	"math"
	"strconv"

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
