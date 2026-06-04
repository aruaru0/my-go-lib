# MaxGain — Maximum Gain in K Moves

## Functions

```go
func MaxGain(N, K int, p, c []int) int
```

MaxGain computes the maximum total gain achievable by moving from 1 to K times along a functional graph starting from node 0. `p[i]` is the 1-indexed destination and `c[i]` is the gain at node i.

Handles cycles by detecting the loop and computing optimal repeats.
