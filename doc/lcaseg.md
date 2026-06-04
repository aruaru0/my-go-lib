# LCASeg — Lowest Common Ancestor (Segment Tree / Euler Tour)

## Types

```go
type LCASeg struct { ... }
```

## Functions

```go
func NewLCASeg(root, n int, node [][]int) *LCASeg
func (l *LCASeg) LCA(u, v int) int
func (l *LCASeg) Dist(u, v int) int
```

LCASeg provides LCA queries using Euler tour + segment tree (RMQ on depth array).

- NewLCASeg creates the structure with the given root, node count, and adjacency list.
- LCA returns the LCA of u and v.
- Dist returns the number of edges between u and v.
