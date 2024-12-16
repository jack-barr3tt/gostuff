package lines

import (
	"github.com/jack-barr3tt/gostuff/nums"
	"github.com/jack-barr3tt/gostuff/types"
)

type Line struct {
	mn int
	md int
	cn int
	cd int
}

func NewFracMC(mn, md, cn, cd int) Line {
	mgcd := nums.Gcd(mn, md)
	cgcd := nums.Gcd(cn, cd)
	if md/mgcd < 0 {
		mn, md = -mn, -md
	}
	if cd/cgcd < 0 {
		cn, cd = -cn, -cd
	}
	return Line{mn / mgcd, md / mgcd, cn / cgcd, cd / cgcd}
}

func NewYMXC[T ~int | ~float64 | ~int64 | ~float32](m, c T) Line {
	mn, md := nums.Rationalize(m, 100000)
	cn, cd := nums.Rationalize(c, 100000)
	return NewFracMC(mn, md, cn, cd)
}

func NewBetween(a, b types.Point) Line {
	mn := a[1] - b[1]
	md := a[0] - b[0]
	cn := a[1]*(a[0]-b[0]) - a[0]*(a[1]-b[1])
	cd := a[0] - b[0]
	return NewFracMC(mn, md, cn, cd)
}

func NewPointDir(p types.Point, d types.Direction) Line {
	mn := d[1]
	md := d[0]
	cn := p[1]*d[0] - p[0]*d[1]
	cd := d[0]
	return NewFracMC(mn, md, cn, cd)
}

func NewAXBYC[T ~int | ~float64 | ~int64 | ~float32](a, b, c T) Line {
	a0, a1 := nums.Rationalize(a, 100000)
	b0, b1 := nums.Rationalize(b, 100000)
	c0, c1 := nums.Rationalize(c, 100000)

	mn := -b1 * a0
	md := a1 * b0
	cn := c0 * b1
	cd := c1 * b0
	return NewFracMC(mn, md, cn, cd)
}

func (l Line) SubX(x float64) float64 {
	return (float64(l.mn)*x)/float64(l.md) + (float64(l.cn) / float64(l.cd))
}

func (l Line) SubY(y float64) float64 {
	return ((y - float64(l.cn)/float64(l.cd)) * float64(l.md)) / float64(l.mn)
}

func (l Line) IntersectsAt(l2 Line) (float64, float64, bool) {
	if l.mn == l2.mn && l.md == l2.md {
		return 0, 0, false
	}
	xn := l.md * l2.md * (l.cn*l2.cd - l.cd*l2.cn)
	xd := l.cd * l2.cd * (l.md*l2.mn - l2.md*l.mn)
	yn := l.cd*l2.cn*l.mn*l2.md - l.cn*l2.cd*l.md*l2.mn
	yd := l.cd * l2.cd * (l.mn*l2.md - l.md*l2.mn)

	println(xn, xd, yn, yd)

	x := float64(xn) / float64(xd)
	if xn%xd == 0 {
		x = float64(xn / xd)
	}
	y := float64(yn) / float64(yd)
	if yn%yd == 0 {
		y = float64(yn / yd)
	}

	return x, y, true
}
