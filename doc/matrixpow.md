# matrixpow

Matrix exponentiation with generic integer types.

## Functions

```go
func PowMatrix[T Integer](A [][]T, p int, mod T) [][]T
```

## Example

```go
A := [][]int{{1, 1}, {1, 0}}
P := mylib.PowMatrix(A, 10, 1000000007)
// P = [[89, 55], [55, 34]]  (Fibonacci)
```
