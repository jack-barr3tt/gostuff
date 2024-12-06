package types

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestEquals(t *testing.T) {
	p1 := Pair[int, string]{First: 1, Second: "a"}
	p2 := Pair[int, string]{First: 1, Second: "a"}
	p3 := Pair[int, string]{First: 2, Second: "a"}

	test.AssertEqual(t, p1.Equals(p2), true)
	test.AssertEqual(t, p1.Equals(p3), false)
	test.AssertEqual(t, p2.Equals(p3), false)
}
