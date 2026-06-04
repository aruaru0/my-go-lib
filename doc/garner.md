# garner

Garner's algorithm for Chinese Remainder Theorem.

## Functions

```go
func Garner(x, m []int, mod int) int
```

Given `x[i] ≡ result (mod m[i])`, computes `result` modulo `mod`. If `mod` is 0, computes the exact integer (may overflow for large values).

## Example

```go
x := []int{2, 3, 2}
m := []int{3, 5, 7}
result := mylib.Garner(x, m, 0) // 23
// 23 ≡ 2 (mod 3), 23 ≡ 3 (mod 5), 23 ≡ 2 (mod 7)
```
