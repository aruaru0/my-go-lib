# DFSRoute

`DFSRoute` finds a path between two nodes in a tree or acyclic graph using depth-first search.

## Types

```go
type Node struct {
    To []int
}
```

## Function

```go
func DFSRoute(nodes []Node, from, to int) []int
```

Returns the path from `from` to `to` as a slice of vertex indices, or `nil` if no path exists. The graph must not contain cycles.

## Example

```go
nodes := []Node{
    {To: []int{1, 2}},
    {To: []int{0, 3}},
    {To: []int{0, 3}},
    {To: []int{1, 2}},
}
route := DFSRoute(nodes, 0, 3)
fmt.Println(route) // [0 1 3] or [0 2 3]
```

## Complexity

O(V + E) time, O(V) space.
