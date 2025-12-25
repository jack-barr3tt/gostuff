package maze

import (
	"testing"

	"github.com/jack-barr3tt/gostuff/test"
	"github.com/jack-barr3tt/gostuff/types"
)

func TestNewMaze(t *testing.T) {
	raw := `######R
#     #
#   # #
#  #  #
# ### #
#     #
L######`

	expected := Maze[rune]{
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

func TestNewBlankMaze(t *testing.T) {
	g1 := NewBlankMaze(3, 3, '.')

	expected1 := Maze[rune]{
		{'.', '.', '.'},
		{'.', '.', '.'},
		{'.', '.', '.'},
	}

	test.AssertEqual(t, g1, expected1)

	g2 := NewBlankMaze(4, 2, '.')

	expected2 := Maze[rune]{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}

	test.AssertEqual(t, g2, expected2)
}

func TestMove(t *testing.T) {
	maze := NewMaze(`######
#    #
#    #
#    #
#    #
######`)

	n, ok := maze.Move(types.Point{1, 1}, types.North)
	test.AssertEqual(t, n, types.Point{1, 2})
	test.AssertEqual(t, ok, true)

	n, ok = maze.Move(types.Point{1, 1}, types.East)
	test.AssertEqual(t, n, types.Point{2, 1})
	test.AssertEqual(t, ok, true)

	n, ok = maze.Move(types.Point{0, 0}, types.South)
	test.AssertEqual(t, n, types.Point{0, 0})
	test.AssertEqual(t, ok, false)
}

func TestAt(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	test.AssertEqual(t, maze.At(types.Point{1, 1}), ' ')
	test.AssertEqual(t, maze.At(types.Point{2, 1}), 'L')
	test.AssertEqual(t, maze.At(types.Point{4, 4}), '*')
}

func TestLocateAll(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	test.AssertEqual(t, maze.LocateAll('*'), []types.Point{{4, 4}})
	test.AssertEqual(t, maze.LocateAll('L'), []types.Point{{2, 1}})
}

func TestSet(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	maze.Set(types.Point{1, 1}, 'X')
	test.AssertEqual(t, maze.At(types.Point{1, 1}), 'X')
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

	clone.Set(types.Point{2, 1}, 'X')

	test.AssertEqual(t, maze.At(types.Point{2, 1}), 'L')
	test.AssertEqual(t, clone.At(types.Point{2, 1}), 'X')
}

func TestPointClone(t *testing.T) {
	p := types.Point{1, 2}
	clone := p.Clone()

	test.AssertEqual(t, p, clone)

	clone[0] = 3

	test.AssertEqual(t, p[0], 1)
	test.AssertEqual(t, clone[0], 3)
}

func TestMazeRotate(t *testing.T) {
	maze := NewMaze(`######
#   *#
#    #
#    #
# L  #
######`)

	rotated := maze.Rotate(0)

	test.AssertEqual(t, maze, rotated)

	expected90 := NewMaze(`######
#*   #
#    #
#   L#
#    #
######`)

	rotated = maze.Rotate(90)
	test.AssertEqual(t, rotated, expected90)

	expected180 := NewMaze(`######
#  L #
#    #
#    #
#*   #
######`)

	rotated = maze.Rotate(180)
	test.AssertEqual(t, rotated, expected180)
}

func TestSubMazeAt(t *testing.T) {
	maze := NewMaze(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	subMaze := NewMaze(`S.S
.A.
M.M`)

	test.AssertEqual(t, maze.SubMazeAt(subMaze, types.Point{0, 1}, []rune{'.'}), true)

	subMaze2 := NewMaze(`M.S
.A.
M.S`)

	test.AssertEqual(t, maze.SubMazeAt(subMaze2, types.Point{0, 1}, []rune{'.'}), false)

	test.AssertEqual(t, maze.SubMazeAt(subMaze2, types.Point{9, 9}, []rune{'.'}), false)
}

func TestFloodFill(t *testing.T) {
	maze := NewMaze(`######
#    #
# ## #
# ## #
#    #
######`)

	expected := NewMaze(`######
#XXXX#
#X##X#
#X##X#
#XXXX#
######`)

	maze.FloodFill(types.Point{1, 1}, ' ', 'X')

	test.AssertEqual(t, maze, expected)
}

func TestInsertSubMazeAt(t *testing.T) {
	maze := NewMaze(`#.#
.#.
#.#`)

	subMaze := NewMaze(`.#.
#.#
.#.`)

	expected := NewMaze(`###
###
###`)

	success := maze.InsertSubMazeAt(subMaze, types.Point{0, 0}, []rune{'.'})

	test.AssertEqual(t, success, true)
	test.AssertEqual(t, maze, expected)
}
