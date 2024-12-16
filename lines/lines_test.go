package lines

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
	"github.com/jack-barr3tt/gostuff/types"
)

func TestNewFracMC(t *testing.T) {
	test.AssertEqual(t, NewFracMC(1, 2, 3, 4), Line{1, 2, 3, 4})
	test.AssertEqual(t, NewFracMC(2, 4, 6, 8), Line{1, 2, 3, 4})
	test.AssertEqual(t, NewFracMC(1, -2, 1, 1), Line{-1, 2, 1, 1})
}

func TestNewYMXC(t *testing.T) {
	test.AssertEqual(t, NewYMXC(1, 2), Line{1, 1, 2, 1})
}

func TestNewBetween(t *testing.T) {
	test.AssertEqual(t, NewBetween(types.Point{1, 2}, types.Point{3, 4}), Line{1, 1, 1, 1})
}

func TestNewPointDir(t *testing.T) {
	test.AssertEqual(t, NewPointDir(types.Point{1, 2}, types.Direction{2, 4}), Line{2, 1, 0, 1})
}

func TestNewAXBYC(t *testing.T) {
	test.AssertEqual(t, NewAXBYC(94, 22, 8400), Line{-47, 11, 4200, 11})
}

func TestSubX(t *testing.T) {
	l := Line{1, 1, 2, 1}
	test.AssertEqual(t, l.SubX(3), 5)
}

func TestSubY(t *testing.T) {
	l := Line{1, 1, 2, 1}
	test.AssertEqual(t, l.SubY(3), 1)
}

func TestIntersectAt(t *testing.T) {
	l1 := Line{1, 1, 2, 1}
	l2 := Line{1, 2, 4, 1}
	x, y, ok := l1.IntersectsAt(l2)
	test.AssertEqual(t, x, 4)
	test.AssertEqual(t, y, 6)
	test.AssertEqual(t, ok, true)

	l3 := Line{1, 1, 5, 1}
	x, y, ok = l1.IntersectsAt(l3)
	test.AssertEqual(t, x, 0)
	test.AssertEqual(t, y, 0)
	test.AssertEqual(t, ok, false)

	l4 := NewAXBYC(94, 22, 8400)
	l5 := NewAXBYC(34, 67, 5400)
	x, y, ok = l4.IntersectsAt(l5)
	test.AssertEqual(t, x, 80)
	test.AssertEqual(t, y, 40)
	test.AssertEqual(t, ok, true)

	l6 := NewAXBYC(26, 67, 10000000012748)
	l7 := NewAXBYC(66, 21, 10000000012176)
	x, y, ok = l6.IntersectsAt(l7)
	test.AssertEqual(t, x, 118679050709)
	test.AssertEqual(t, y, 103199174542)
	test.AssertEqual(t, ok, true)
}
