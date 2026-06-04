# LazySegTree — ジェネリック遅延評価セグメント木

遅延伝搬に対応したセグメント木です。範囲更新と範囲クエリをサポートします。

## 関数

- `NewLazySegTree[S, F any](v []S, e func() S, merger func(a, b S) S, mapper func(f F, x S) S, comp func(f, g F) F, id func() F) *LazySegTree[S, F]` — 木を作成します
  - `e`: S の単位元
  - `merger`: 2 つの S 値を結合します
  - `mapper`: F を S に適用します
  - `comp`: 2 つの F 値を合成します
  - `id`: F の単位元
- `(*LazySegTree[S, F]).Set(p int, x S)` — 点更新
- `(*LazySegTree[S, F]).Get(p int) S` — 点取得
- `(*LazySegTree[S, F]).Prod(l, r int) S` — `[l, r)` の範囲クエリ
- `(*LazySegTree[S, F]).AllProd() S` — 全要素のクエリ
- `(*LazySegTree[S, F]).Apply(p int, f F)` — 点適用
- `(*LazySegTree[S, F]).RangeApply(l, r int, f F)` — `[l, r)` の範囲適用
- `(*LazySegTree[S, F]).MaxRight(l int, cmp func(S) bool) int` — 右方向への二分探索
- `(*LazySegTree[S, F]).MinLeft(r int, cmp func(S) bool) int` — 左方向への二分探索
