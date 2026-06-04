# mathutil

Mathematical utility functions.

## Functions

```go
func FloorSum(n, m, a, b int) int
func Crt(r, m []int) [2]int
func PowMod(x, n, m int) int
func SafeMod(x, d int) int
func InvMod(x, m int) int
func InvGcd(a, b int) [2]int
func PrimitiveRoot(m int) int
```

## FloorSum

Computes `sum_{i=0}^{n-1} floor((a*i + b) / m)`.

## Crt

Chinese Remainder Theorem. Returns `[r, m]` such that `x ≡ r[i] (mod m[i])`. Returns `[0, 0]` if no solution.

## PowMod

Binary exponentiation: `x^n mod m`.

## SafeMod

Returns `x mod d` ensuring non-negative result.

## InvMod

Modular inverse of `x` modulo `m` (requires `gcd(x, m) = 1`).

## InvGcd

Returns `[g, inv]` where `g = gcd(a, b)` and `inv` is the modular inverse of `a/g` modulo `b/g`.

## PrimitiveRoot

Finds the smallest primitive root modulo `m`.
