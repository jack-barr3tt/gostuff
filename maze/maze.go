package maze

import (
	"strings"

	"github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

type Maze [][]rune

func NewMaze(raw string) Maze {
	lines := strings.Split(raw, "\n")

	maze := make([][]rune, len(lines))

	for i, line := range lines {
		maze[len(lines)-1-i] = []rune(line)
	}

	return maze
}

func NewBlankMaze(width, height int) Maze {
	maze := make([][]rune, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			maze[y] = append(maze[y], '.')
		}
	}
	return maze
}

func (m Maze) Move(p types.Point, d types.Direction) (types.Point, bool) {
	newPos := types.Point{p[0] + d[0], p[1] + d[1]}

	if newPos[1] < 0 || newPos[1] >= len(m) || newPos[0] < 0 || newPos[0] >= len(m[0]) {
		return p, false
	}

	return newPos, true
}

func (m Maze) At(p types.Point) rune {
	return m[p[1]][p[0]]
}

func (m Maze) LocateAll(r rune) []types.Point {
	points := []types.Point{}

	for i, row := range m {
		for j, cell := range row {
			if cell == r {
				points = append(points, types.Point{j, i})
			}
		}
	}

	return points
}

func (m Maze) Set(p types.Point, r rune) {
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

func (m Maze) Print() {
	output := ""
	for i := range m {
		for j := range m[len(m)-1-i] {
			output += string(m[len(m)-1-i][j])
		}
		output += "\n"
	}
	println(output)
}

func (m Maze) Rotate(deg int) Maze {
	normalised := ((deg % 360) + 360) % 360
	if normalised%90 != 0 {
		panic("Can only rotate in 90 degree increments")
	}
	if len(m) != len(m[0]) {
		panic("Can only rotate square mazes")
	}

	rotated := m.Clone()

	for i := 0; i < normalised/90; i++ {
		temp := rotated.Clone()
		for y := range rotated {
			for x := range rotated[y] {
				temp[x][len(rotated)-1-y] = rotated[y][x]
			}
		}
		rotated = temp
	}

	return rotated
}

func (m Maze) SubMazeAt(m2 Maze, origin types.Point, ignore []rune) bool {
	ignoreMap := slices.Frequency(ignore)
	for i, row := range m2 {
		for j, cell := range row {
			if _, ok := ignoreMap[cell]; ok {
				continue
			}

			p, ok := m.Move(origin, types.North.Multiply(i).Add(types.East.Multiply(j)))
			if !ok || m.At(p) != cell {
				return false
			}
		}
	}

	return true
}
