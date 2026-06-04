# ConvexHull — Monotone Chain

## Types

```go
type Point struct {
	X, Y int
}
```

## Functions

```go
func ConvexHull(pts []Point) []Point
```

ConvexHull returns the convex hull of a set of points in CCW order starting from the leftmost point, using the Monotone Chain (Andrew's) algorithm.

All points are collinear: returns 2 endpoints. Single point: returns that point.
