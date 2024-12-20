package types

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestRotate(t *testing.T) {
	test.AssertEqual(t, North.Rotate(90), East)
	test.AssertEqual(t, North.Rotate(-90), West)
	test.AssertEqual(t, North.Rotate(45), NorthEast)
	test.AssertEqual(t, North.Rotate(-45), NorthWest)

	test.AssertEqual(t, East.Rotate(90), South)
	test.AssertEqual(t, East.Rotate(-90), North)
	test.AssertEqual(t, East.Rotate(45), SouthEast)
	test.AssertEqual(t, East.Rotate(-45), NorthEast)
}
func TestInverse(t *testing.T) {
	// check standard compass directions
	test.AssertEqual(t, North.Inverse(), South)
	test.AssertEqual(t, East.Inverse(), West)
	test.AssertEqual(t, South.Inverse(), North)
	test.AssertEqual(t, West.Inverse(), East)

	// test ad-hoc directions
	test.AssertEqual(t, Direction{1, -2}.Inverse(), Direction{-1, 2})
}

func TestMultiply(t *testing.T) {
	// check standard compass directions
	test.AssertEqual(t, North.Multiply(2), Direction{0, 2})
	test.AssertEqual(t, East.Multiply(3), Direction{3, 0})
	test.AssertEqual(t, South.Multiply(4), Direction{0, -4})
	test.AssertEqual(t, West.Multiply(5), Direction{-5, 0})

	// test ad-hoc directions
	test.AssertEqual(t, Direction{1, -2}.Multiply(2), Direction{2, -4})
}

func TestAdd(t *testing.T) {
	// check standard compass directions
	test.AssertEqual(t, North.Add(East), NorthEast)
	test.AssertEqual(t, East.Add(South), SouthEast)
	test.AssertEqual(t, South.Add(West), SouthWest)
	test.AssertEqual(t, West.Add(North), NorthWest)

	// test ad-hoc directions
	test.AssertEqual(t, Direction{1, -2}.Add(Direction{3, 4}), Direction{4, 2})
}

func TestManhattan(t *testing.T) {
	test.AssertEqual(t, North.Manhattan(), 1)
	test.AssertEqual(t, Direction{1, -2}.Manhattan(), 3)
}

func TestDirFromSlice(t *testing.T) {
	test.AssertEqual(t, Direction{1, 2}, DirFromSlice([]int{1, 2}))
}
