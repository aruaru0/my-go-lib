# ConvexHullTrick — Li Chao Segment Tree

Maintains a set of lines `y = a*x + b` and queries the maximum value at a given `x`.

## Functions

- `NewConvexHullTrick[T Ordered](pos []T) *ConvexHullTrick[T]` — creates a tree over x-coordinates `pos` (sorted, deduped internally)
- `(*ConvexHullTrick[T]).AddLine(a, b T)` — adds line `a*x + b`
- `(*ConvexHullTrick[T]).GetMax(x T) T` — returns max value among all lines at x
