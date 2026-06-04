# gacha

Gacha (coupon collector) completion expectation.

## Functions

```go
func GachaExpectation(n int) float64
```

## GachaExpectation

Returns the expected number of draws needed to collect all `n` distinct items
in a coupon collector problem. Computed as `n * H_n` where `H_n` is the n-th
harmonic number.
