# SegTree — ジェネリックセグメント木

任意の型 `T` と任意のモノイド演算に対応したジェネリックセグメント木です。

## 関数

- `NewSegTree[T any](v []T, e func() T, m func(a, b T) T) *SegTree[T]` — 初期スライス `v`、単位元 `e()`、結合演算 `m` から木を作成します
- `(*SegTree[T]).Set(p int, x T)` — インデックス `p` の要素を `x` に設定します
- `(*SegTree[T]).Get(p int) T` — インデックス `p` の要素を取得します
- `(*SegTree[T]).Prod(l, r int) T` — `[l, r)` の `m`-積を返します
- `(*SegTree[T]).AllProd() T` — 全要素の `m`-積を返します
- `(*SegTree[T]).MaxRight(l int, cmp func(T) bool) int` — `Prod(l, r)` が `cmp` を満たす最大の `r` を返します
- `(*SegTree[T]).MinLeft(r int, cmp func(T) bool) int` — `Prod(l, r)` が `cmp` を満たす最小の `l` を返します
