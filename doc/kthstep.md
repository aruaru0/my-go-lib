# KthStep — K-Step Graph Traversal by Doubling

## Functions

```go
func KthStep(N int, K int64, next []int) int
```

KthStep returns the node reached after exactly K steps starting from node 0, where `next[i]` is the destination from node i in one step.

Uses binary lifting (doubling) with a 60-bit table.
