package mylib

import "sort"

type Point struct {
	X, Y int
}

type points []Point

func (p points) Len() int      { return len(p) }
func (p points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p points) Less(i, j int) bool {
	if p[i].X == p[j].X {
		return p[i].Y < p[j].Y
	}
	return p[i].X < p[j].X
}

func ccw(a, b, c Point) bool {
	return ((b.X-a.X)*(c.Y-a.Y)) > ((b.Y-a.Y)*(c.X-a.X))
}

func ConvexHull(pts []Point) []Point {
	if len(pts) < 2 {
		return pts
	}
	sort.Sort(points(pts))
	var h []Point

	for _, pt := range pts {
		for len(h) >= 2 && !ccw(h[len(h)-2], h[len(h)-1], pt) {
			h = h[:len(h)-1]
		}
		h = append(h, pt)
	}

	t := len(h) + 1
	for i := len(pts) - 2; i >= 0; i-- {
		pt := pts[i]
		for len(h) >= t && !ccw(h[len(h)-2], h[len(h)-1], pt) {
			h = h[:len(h)-1]
		}
		h = append(h, pt)
	}

	return h[:len(h)-1]
}
