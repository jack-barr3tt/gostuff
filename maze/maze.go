package maze

import (
	"strings"

	"github.com/jack-barr3tt/gostuff/slices"
)

type Maze [][]rune

type Point [2]int

type Direction [2]int

var (
	North     = Direction{0, 1}
	East      = Direction{1, 0}
	South     = Direction{0, -1}
	West      = Direction{-1, 0}
	NorthEast = Direction{1, 1}
	SouthEast = Direction{1, -1}
	NorthWest = Direction{-1, 1}
	SouthWest = Direction{-1, -1}
)
var step90 = []Direction{North, East, South, West}
var step45 = []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}

type Rotation string

var (
	C45  = Rotation("45C")
	C90  = Rotation("90C")
	CC45 = Rotation("45CC")
	CC90 = Rotation("90CC")
)

func NewMaze(raw string) Maze {
	lines := strings.Split(raw, "\n")

	maze := make([][]rune, len(lines))

	for i, line := range lines {
		maze[len(lines)-1-i] = []rune(line)
	}

	return maze
}

func (d Direction) RotateDirection(r Rotation) Direction {
	if r == C90 || r == CC90 {
		i := slices.FindIndex(func(x Direction) bool { return x == d }, step90)

		if r == C90 {
			return step90[(i+1)%4]
		} else {
			return step90[(i+3)%4]
		}
	}
	if r == C45 || r == CC45 {
		i := slices.FindIndex(func(x Direction) bool { return x == d }, step45)
		if r == C45 {
			return step45[(i+1)%8]
		} else {
			return step45[(i+7)%8]
		}
	}

	return d
}

func (d Direction) Inverse() Direction {
	return Direction{-d[0], -d[1]}
}

func (m Maze) Move(p Point, d Direction) (Point, bool) {
	newPos := Point{p[0] + d[0], p[1] + d[1]}

	if newPos[1] < 0 || newPos[1] >= len(m) || newPos[0] < 0 || newPos[0] >= len(m[0]) {
		return p, false
	}

	return newPos, true
}

func (m Maze) At(p Point) rune {
	return m[p[1]][p[0]]
}

func (m Maze) LocateAll(r rune) []Point {
	points := []Point{}

	for i, row := range m {
		for j, cell := range row {
			if cell == r {
				points = append(points, Point{j, i})
			}
		}
	}

	return points
}

func (m Maze) Set(p Point, r rune) {
	m[p[1]][p[0]] = r
}

func (m Maze) Clone() Maze {
	clone := make([][]rune, len(m))

	for i, row := range m {
		clone[i] = make([]rune, len(row))
		copy(clone[i], row)
	}

	return clone
}

func (p Point) Clone() Point {
	return Point{p[0], p[1]}
}

func DirectionBetween(a, b Point) Direction {
	return Direction{b[0] - a[0], b[1] - a[1]}
}
