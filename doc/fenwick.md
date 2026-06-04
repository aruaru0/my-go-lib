# FenwickTree — Generic Fenwick Tree (Binary Indexed Tree)

Generic implementation supporting any `Integer` type (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`).

## Functions

- `NewFenwickTree[T Integer](n int) *FenwickTree[T]` — creates a tree of size n
- `(*FenwickTree[T]).Add(i int, x T)` — adds x to element at index i
- `(*FenwickTree[T]).Sum(i int) T` — returns sum of elements in [0, i] (inclusive). Returns 0 for negative i.
- `(*FenwickTree[T]).RangeSum(l, r int) T` — returns sum of elements in [l, r)
- `(*FenwickTree[T]).LowerBound(x T) int` — returns smallest index such that Sum(index) >= x
