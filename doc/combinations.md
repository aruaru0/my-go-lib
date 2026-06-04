# combinations

Combination enumeration utilities.

## Functions

```go
func Combinations[T any](list []T, choose int) <-chan []T
func CombinationsRec(n, k int) <-chan []int
```

## Combinations

Generates all combinations of `choose` elements from `list`, streaming them
through a channel using goroutines.

## CombinationsRec

Generates all index combinations (indices 0 to n-1 choose k) through a channel
using recursive backtracking.
