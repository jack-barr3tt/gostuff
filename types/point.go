package types

type Point [2]int

func (p Point) UnsafeMove(d Direction) Point {
	return Point{p[0] + d[0], p[1] + d[1]}
}

func (p Point) Clone() Point {
	return Point{p[0], p[1]}
}

func (a Point) DirectionTo(b Point) Direction {
	return Direction{b[0] - a[0], b[1] - a[1]}
}

func PointFromSlice(s []int) Point {
	return Point{s[0], s[1]}
}
