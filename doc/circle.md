# Circle — Circle-Circle Intersection

## Types

```go
type PointF struct {
	X, Y float64
}
```

## Functions

```go
func CircleCross(p0, p1 PointF, r1, r2 float64) []PointF
```

CircleCross returns the intersection points of two circles. Returns nil if the circles do not intersect (separate or concentric).
