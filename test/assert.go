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

func AssertEqual[A any](t *testing.T, a A, b A) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %v, got %v", a, b)
	}
}

func AssertCases[A, B any](t *testing.T, fn func(A) B, cases []Case[A, B]) {
	for _, c := range cases {
		fmt.Println(c.Input, c.Expected, fn(c.Input))
		AssertEqual(t, c.Expected, fn(c.Input))
	}
}
