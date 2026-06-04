package mylib

import "math"

type PointF struct {
	X, Y float64
}

func CircleCross(p0, p1 PointF, r1, r2 float64) []PointF {
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y
	l2 := dx*dx + dy*dy
	a := (l2 + r1*r1 - r2*r2) / 2.0
	D := l2*r1*r1 - a*a
	if D < 0 || l2 == 0 {
		return nil
	}
	D = math.Sqrt(D)
	return []PointF{
		{(a*dx + D*dy) / l2 + p0.X, (a*dy - D*dx) / l2 + p0.Y},
		{(a*dx - D*dy) / l2 + p0.X, (a*dy + D*dx) / l2 + p0.Y},
	}
}
