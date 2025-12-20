package types

import (
	"math"

	"github.com/jack-barr3tt/gostuff/nums"
)

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

func (d Direction) Rotate(deg int) Direction {
	rad := float64(deg) * math.Pi / 180.0

	ct := math.Cos(rad)
	st := math.Sin(rad)

	newX := float64(d[0])*ct + float64(d[1])*st
	newY := -float64(d[0])*st + float64(d[1])*ct

	return Direction{int(math.Round(newX)), int(math.Round(newY))}
}

func (d Direction) Inverse() Direction {
	return d.Multiply(-1)
}

func (d Direction) Multiply(n int) Direction {
	return Direction{d[0] * n, d[1] * n}
}

func (d Direction) Add(other Direction) Direction {
	return Direction{d[0] + other[0], d[1] + other[1]}
}

func (d Direction) Manhattan() int {
	return nums.Abs(d[0]) + nums.Abs(d[1])
}

func (d Direction) Magnitude() float64 {
	return math.Sqrt(float64(d[0]*d[0] + d[1]*d[1]))
}

func (d Direction) Unit() Direction {
	return Direction{d[0] / int(d.Magnitude()), d[1] / int(d.Magnitude())}
}

func DirFromSlice(s []int) Direction {
	return Direction{s[0], s[1]}
}
