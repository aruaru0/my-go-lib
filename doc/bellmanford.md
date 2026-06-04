# BellmanFord

`BellmanFord` computes shortest paths from a single source in a weighted directed graph, handling negative edges and detecting negative cycles reachable from the source.

## Types

```go
type Edge struct {
    From, To, Cost int
}
```

## Function

```go
func BellmanFord(N, S, E int, edges []Edge) ([]int, bool)
```

- `N`: number of vertices (0..N-1)
- `S`: source vertex
- `E`: target vertex (used only for negative cycle check)
- `edges`: list of directed edges
- Returns: distance array, and a bool indicating whether `E` is reachable from a negative cycle

## Example

```go
edges := []Edge{
    {0, 1, 5},
    {0, 2, 3},
    {1, 3, 2},
    {2, 3, 4},
}
dist, neg := BellmanFord(4, 0, 3, edges)
fmt.Println(dist[3]) // 7
fmt.Println(neg)     // false
```

## Complexity

O(V·E) time, O(V) space.
