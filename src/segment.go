package mylib

func SegmentsIntersect(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	ta := (cx-dx)*(ay-cy) + (cy-dy)*(cx-ax)
	tb := (cx-dx)*(by-cy) + (cy-dy)*(cx-bx)
	tc := (ax-bx)*(cy-ay) + (ay-by)*(ax-cx)
	td := (ax-bx)*(dy-ay) + (ay-by)*(ax-dx)

	return tc*td < 0 && ta*tb < 0
}
