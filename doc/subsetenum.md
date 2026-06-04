# SubsetEnum — Subset Enumeration of Set Bits

## Functions

```go
func Subsets(i int) []int
```

Subsets returns all non-zero submasks of `i`. For example, `Subsets(0b1110)` returns `[0b1110, 0b1100, 0b1010, 0b1000, 0b0110, 0b0100, 0b0010]`.

Uses the standard submask enumeration pattern `for j := i; j != 0; j = (j - 1) & i`.
