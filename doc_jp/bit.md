# BIT — Binary Indexed Tree (int 用 Fenwick 木)

## 関数

- `NewBIT(n int) *BIT` — サイズ n の BIT を作成します (0-indexed の位置 0..n-1)
- `(*BIT).Add(k, v int)` — インデックス k の要素に v を加算します
- `(*BIT).Sum(k int) int` — [0, k] (両端を含む) の要素の和を返します。k < 0 の場合は 0、k >= n の場合は n-1 にクランプされます。
- `(*BIT).RangeSum(l, r int) int` — [l, r) の要素の和を返します (l を含み、r は含みません)。l >= r の場合は 0 を返します。
