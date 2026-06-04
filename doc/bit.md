# BIT — Binary Indexed Tree (Fenwick Tree for int)

## Functions

- `NewBIT(n int) *BIT` — creates a BIT of size n (0-indexed positions 0..n-1)
- `(*BIT).Add(k, v int)` — adds v to element at index k
- `(*BIT).Sum(k int) int` — returns sum of elements in [0, k] (inclusive). Returns 0 if k < 0, clamps to n-1 if k >= n.
- `(*BIT).RangeSum(l, r int) int` — returns sum of elements in [l, r) (l inclusive, r exclusive). Returns 0 if l >= r.
