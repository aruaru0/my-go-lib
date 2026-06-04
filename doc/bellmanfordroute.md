# BellmanFordRoute

`BellmanFordRoute` extends Bellman-Ford to also track the actual route (path) from the source to each vertex.

## Types

```go
type Edge struct {
    From, To, Cost int
}

type Route struct {
    Path []int
}
```

## Function

```go
func BellmanFordRoute(N, S int, edges []Edge) ([]int, []Route, bool)
```

- `N`: number of vertices
- `S`: source vertex
- `edges`: list of directed edges
- Returns: distance array, route array (one per vertex), and a bool indicating a negative cycle

## Example

```go
edges := []Edge{{0, 1, 2}, {1, 2, 3}, {0, 2, 6}}
dist, routes, neg := BellmanFordRoute(3, 0, edges)
fmt.Println(dist[2])       // 5
fmt.Println(routes[2].Path) // [0 1 2]
```

## Complexity

O(V·E) time, O(V·E) space in worst case for routes.
