# FenwickTree — ジェネリック Fenwick 木 (Binary Indexed Tree)

任意の `Integer` 型 (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`) をサポートするジェネリック実装です。

## 関数

- `NewFenwickTree[T Integer](n int) *FenwickTree[T]` — サイズ n の木を作成します
- `(*FenwickTree[T]).Add(i int, x T)` — インデックス i の要素に x を加算します
- `(*FenwickTree[T]).Sum(i int) T` — [0, i] (両端を含む) の要素の和を返します。負の i の場合は 0 を返します。
- `(*FenwickTree[T]).RangeSum(l, r int) T` — [l, r) の要素の和を返します
- `(*FenwickTree[T]).LowerBound(x T) int` — `Sum(index) >= x` となる最小のインデックスを返します
