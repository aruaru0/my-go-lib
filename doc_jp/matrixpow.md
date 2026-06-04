# matrixpow

ジェネリック整数型に対応した行列の累乗です。

## 関数

```go
func PowMatrix[T Integer](A [][]T, p int, mod T) [][]T
```

## 使用例

```go
A := [][]int{{1, 1}, {1, 0}}
P := mylib.PowMatrix(A, 10, 1000000007)
// P = [[89, 55], [55, 34]]  (フィボナッチ数)
```
