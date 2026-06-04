# cyclesum

Efficient sum computation for sequences with cycles using modulo-based cycle detection.

## Functions

```go
func CycleSum(start, N, mod int, next func(int) int) int
```

## CycleSum

Computes the sum of the first `N` terms of a sequence defined by `next`,
starting from `start`. Detects cycles based on repeated modulo residues and
uses the cycle to compute the sum in O(mod) time regardless of N.
