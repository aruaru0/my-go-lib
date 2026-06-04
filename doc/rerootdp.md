# RerootDP — Tree Distances via Rerooting DP

## Functions

```go
func TreeDistances(N int, node [][]int) []int
func TreeDistancesWithNode(N int, node [][]int) ([]int, []int)
```

TreeDistances computes the farthest distance from each node to any other node in an unweighted tree (adjacency list).

TreeDistancesWithNode additionally returns the index of a farthest node (largest index on tie).

Rerooting (two-pass DFS) technique: first DFS computes subtree heights, second DFS propagates parent-side info to children.
