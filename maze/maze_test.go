package maze

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
)

func TestNewMaze(t *testing.T) {
	raw := `######R
#     #
#   # #
#  #  #
# ### #
#     #
L######`

	expected := Maze{
		{'L', '#', '#', '#', '#', '#', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', ' ', '#', '#', '#', ' ', '#'},
		{'#', ' ', ' ', '#', ' ', ' ', '#'},
		{'#', ' ', ' ', ' ', '#', ' ', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', '#', '#', '#', '#', '#', 'R'},
	}

	test.AssertEqual(t, NewMaze(raw), expected)
}

func TestRotateDirection(t *testing.T) {
	test.AssertEqual(t, North.RotateDirection(C90), East)
	test.AssertEqual(t, North.RotateDirection(CC90), West)
	test.AssertEqual(t, North.RotateDirection(C45), NorthEast)
	test.AssertEqual(t, North.RotateDirection(CC45), NorthWest)

	test.AssertEqual(t, East.RotateDirection(C90), South)
	test.AssertEqual(t, East.RotateDirection(CC90), North)
	test.AssertEqual(t, East.RotateDirection(C45), SouthEast)
	test.AssertEqual(t, East.RotateDirection(CC45), NorthEast)
}

func TestMove(t *testing.T) {
	maze := NewMaze(`######
#    #
#    #
#    #
#    #
######`)

	n, ok := maze.Move(Point{1, 1}, North)
	test.AssertEqual(t, n, Point{1, 2})
	test.AssertEqual(t, ok, true)

	n, ok = maze.Move(Point{1, 1}, East)
	test.AssertEqual(t, n, Point{2, 1})
	test.AssertEqual(t, ok, true)

	n, ok = maze.Move(Point{0, 0}, South)
	test.AssertEqual(t, n, Point{0, 0})
	test.AssertEqual(t, ok, false)
}

func TestAt(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	test.AssertEqual(t, maze.At(Point{1, 1}), ' ')
	test.AssertEqual(t, maze.At(Point{2, 1}), 'L')
	test.AssertEqual(t, maze.At(Point{4, 4}), '*')
}

func TestLocateAll(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	test.AssertEqual(t, maze.LocateAll('*'), []Point{{4, 4}})
	test.AssertEqual(t, maze.LocateAll('L'), []Point{{2, 1}})
}

func TestSet(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	maze.Set(Point{1, 1}, 'X')
	test.AssertEqual(t, maze.At(Point{1, 1}), 'X')
}

func TestMazeClone(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	clone := maze.Clone()
	test.AssertEqual(t, maze, clone)

	clone.Set(Point{2, 1}, 'X')

	test.AssertEqual(t, maze.At(Point{2, 1}), 'L')
	test.AssertEqual(t, clone.At(Point{2, 1}), 'X')
}

func TestPointClone(t *testing.T) {
	p := Point{1, 2}
	clone := p.Clone()

	test.AssertEqual(t, p, clone)

	clone[0] = 3

	test.AssertEqual(t, p[0], 1)
	test.AssertEqual(t, clone[0], 3)
}
