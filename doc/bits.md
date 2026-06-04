# bits

Bit manipulation utilities.

## Functions

```go
func CeilPow2(n int) int
func Bsf(n uint) int
```

## CeilPow2

Returns the smallest `x` such that `2^x >= n`.

## Bsf

Bit scan forward. Returns the number of trailing zero bits in `n` (using `bits.TrailingZeros`).
