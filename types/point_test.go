package types

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestPointUnsafeMove(t *testing.T) {
	test.AssertEqual(t, Point{1, 2}.UnsafeMove(Direction{2, 3}), Point{3, 5})
}

func TestDirectionTo(t *testing.T) {
	// check standard compass directions
	test.AssertEqual(t, Point{0, 0}.DirectionTo(Point{1, 0}), East)
	test.AssertEqual(t, Point{0, 0}.DirectionTo(Point{0, 1}), North)
	test.AssertEqual(t, Point{0, 0}.DirectionTo(Point{-1, 0}), West)
	test.AssertEqual(t, Point{0, 0}.DirectionTo(Point{0, -1}), South)

	// test ad-hoc directions
	test.AssertEqual(t, Point{4, 6}.DirectionTo(Point{5, 4}), Direction{1, -2})
	test.AssertEqual(t, Point{4, 6}.DirectionTo(Point{2, 5}), Direction{-2, -1})
}
