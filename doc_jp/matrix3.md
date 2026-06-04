# Matrix3 — 行列の乗算

## 関数

- `Mul3x3(A, B [3][3]int) [3][3]int` — 2 つの 3x3 int 行列を乗算します
- `MulMatrix[T Integer](A, B [][]T) [][]T` — ジェネリック行列乗算です。A が n×m、B が m×p の場合、結果は n×p です。
