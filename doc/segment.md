# Segment — Line Segment Intersection

## Functions

```go
func SegmentsIntersect(ax, ay, bx, by, cx, cy, dx, dy int) bool
```

SegmentsIntersect returns true if segment AB and segment CD intersect strictly (endpoint contact does not count).

Uses the cross product sign test.
