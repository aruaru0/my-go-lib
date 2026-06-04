# xorsum

Sum of XOR over all unordered pairs in an array.

## Functions

```go
func XorSum(dist []int) int64
```

## XorSum

Computes the sum of `dist[i] ^ dist[j]` over all pairs `(i, j)` with `i < j`.
Uses bit-by-bit calculation: for each bit, counts ones and zeros to compute
the contribution `ones * zeros * 2^bit`.
