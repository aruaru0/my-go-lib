# maxgaink

Maximum gain achievable with up to K moves in a functional graph.

## Functions

```go
func MaxGainK(i, N, K int, p, c []int) int
```

## MaxGainK

Starting from position `i`, you can move along edges defined by `p` (1-indexed
next positions) up to `K` times. Each node `v` contributes `c[v]` when visited.
Returns the maximum total gain achievable. Detects cycles in the functional
graph and uses them to accelerate the computation.
