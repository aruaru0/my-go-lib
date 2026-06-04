# SccGraph (Strongly Connected Components)

`SccGraph` computes strongly connected components of a directed graph using Tarjan's algorithm.

## Constructor

```go
func NewSccGraph(n int) *SccGraph
```

Creates a new graph with `n` vertices (0..n-1).

## Methods

```go
func (scc *SccGraph) AddEdge(from, to int)
```
Adds a directed edge from `from` to `to`.

```go
func (scc *SccGraph) Scc() [][]int
```
Returns the list of strongly connected components. Each component is a slice of vertex indices. Components are returned in topological order (reverse DAG order).

## Example

```go
g := NewSccGraph(4)
g.AddEdge(0, 1)
g.AddEdge(1, 2)
g.AddEdge(2, 0)
g.AddEdge(1, 3)
scc := g.Scc()
fmt.Println(scc) // [[3] [0 1 2]]
```

## Complexity

O(V + E) time, O(V + E) space.
