# KthStepLoop — K-Step Graph Traversal with Loop Detection

## Functions

```go
func KthStepLoop(N int, K int64, next []int) int
```

KthStepLoop returns the node reached after exactly K steps starting from node 0. Unlike KthStep, it uses map-based loop detection for efficiency when K is very large and the graph contains cycles.
