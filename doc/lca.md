# LCA — Lowest Common Ancestor (Binary Lifting)

## Types

```go
type LCA struct { ... }
```

## Functions

```go
func NewLCA(n int) *LCA
func (l *LCA) AddEdge(from, to int)
func (l *LCA) Build(root int)
func (l *LCA) LCA(u, v int) int
func (l *LCA) Dist(u, v int) int
```

LCA provides lowest common ancestor queries on a rooted tree using binary lifting (doubling).

- NewLCA creates a tree with n nodes.
- AddEdge adds an undirected edge.
- Build precomputes ancestors (must call after adding all edges, before queries).
- LCA returns the LCA of u and v.
- Dist returns the number of edges between u and v.
