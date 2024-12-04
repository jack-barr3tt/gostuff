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

func AssertCases[A, B any](t *testing.T, fn func(A) B, cases []Case[A, B]) {
	for _, c := range cases {
		fmt.Println(c.Input, c.Expected, fn(c.Input))
		AssertEqual(t, fn(c.Input), c.Expected)
	}
}
