package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Case[A, B any] struct {
	Input    A
	Expected B
}

func AssertEqual[A any](t *testing.T, actual A, expected A) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func AssertNotEqual[A any](t *testing.T, actual A, expected A) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected not to equal %v, got %v", expected, actual)
	}
}

func AssertCases[A, B any](t *testing.T, fn func(A) B, cases []Case[A, B]) {
	for _, c := range cases {
		fmt.Println(c.Input, c.Expected, fn(c.Input))
		AssertEqual(t, fn(c.Input), c.Expected)
	}
}

func AssertSlicesEqual[A any](t *testing.T, actual []A, expected []A) {
	if len(actual) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(actual))
		return
	}

	counts := make(map[string]int)

	for _, exp := range expected {
		key := fmt.Sprintf("%v", exp)
		counts[key]++
	}

	for _, act := range actual {
		key := fmt.Sprintf("%v", act)
		counts[key]--
	}

	for key, count := range counts {
		if count != 0 {
			t.Errorf("Slices differ: element %v has count mismatch", key)
			return
		}
	}
}
