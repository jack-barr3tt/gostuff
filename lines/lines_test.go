package lines

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
	"github.com/jack-barr3tt/gostuff/types"
)

func TestNewYMXC(t *testing.T) {
	test.AssertEqual(t, NewYMXC(1, 2), Line{1, 2})
}

func TestNewBetween(t *testing.T) {
	test.AssertEqual(t, NewBetween(types.Point{1, 2}, types.Point{3, 4}), Line{1, 1})
}

func TestNewPointDir(t *testing.T) {
	test.AssertEqual(t, NewPointDir(types.Point{1, 2}, types.Direction{2, 4}), Line{2, 0})
}

func TestSubX(t *testing.T) {
	l := Line{1, 2}
	test.AssertEqual(t, l.SubX(3), 5)
}

func TestSubY(t *testing.T) {
	l := Line{1, 2}
	test.AssertEqual(t, l.SubY(3), 1)
}

func TestIntersectAt(t *testing.T) {
	l1 := Line{1, 2}
	l2 := Line{0.5, 4}
	x, y, ok := l1.IntersectsAt(l2)
	test.AssertEqual(t, x, 4)
	test.AssertEqual(t, y, 6)
	test.AssertEqual(t, ok, true)

	l3 := Line{1, 5}
	x, y, ok = l1.IntersectsAt(l3)
	test.AssertEqual(t, x, 0)
	test.AssertEqual(t, y, 0)
	test.AssertEqual(t, ok, false)
}
