# extgcd

Extended Euclidean algorithm with generics.

## Functions

```go
func ExtGcd[T Integer](a, b T) (T, T, T)
```

Returns `(g, x, y)` such that `a*x + b*y = g = gcd(a, b)`.

## Example

```go
g, x, y := mylib.ExtGcd(6, 15)
// g = 3, x = -2, y = 1
// 6*(-2) + 15*1 = 3 = gcd(6, 15)
```
