package slices

import (
	"math"
	"strconv"
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
func Reduce[I any, O any](fn func(I, O) O, l []I, init O) O {
	out := init
	for _, v := range l {
		out = fn(v, out)
	}
	return out
}

// Function to convert a slice of strings to a slice of ints
func StrsToInts(l []string) []int {
	return Map(func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}, l)
}

// Zip together two slices
func Zip[A any, B any](a []A, b []B) [][2]any {
	length := int(math.Min(float64(len(a)), float64(len(b))))
	out := make([][2]any, length)
	for i := 0; i < length; i++ {
		out[i] = [2]any{a[i], b[i]}
	}
	return out
}
