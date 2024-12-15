package lines

import "github.com/jack-barr3tt/gostuff/types"

type Line struct {
	m float64
	c float64
}

func NewYMXC[T ~int | ~float64 | ~int64 | ~float32](m, c T) Line {
	return Line{float64(m), float64(c)}
}

func NewBetween(a, b types.Point) Line {
	m := (float64(b[1]) - float64(a[1])) / (float64(b[0]) - float64(a[0]))
	c := float64(a[1]) - m*float64(a[0])
	return Line{m, c}
}

func NewPointDir(p types.Point, d types.Direction) Line {
	m := float64(d[1]) / float64(d[0])
	c := float64(p[1]) - m*float64(p[0])
	return Line{m, c}
}

func NewAXBYC[T ~int | ~float64 | ~int64 | ~float32](a, b, c T) Line {
	return Line{float64(-a) / float64(b), float64(c) / float64(b)}
}

func (l Line) SubX(x float64) float64 {
	return l.m*x + l.c
}

func (l Line) SubY(y float64) float64 {
	return (y - l.c) / l.m
}

func (l Line) IntersectsAt(l2 Line) (float64, float64, bool) {
	if l.m == l2.m {
		return 0, 0, false
	}
	x := (l2.c - l.c) / (l.m - l2.m)
	y := l.SubX(x)
	return x, y, true
}
