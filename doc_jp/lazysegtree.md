# LazySegTree — ジェネリック遅延評価セグメント木

配列の「区間に対する更新（または値の代入・加算など）」と「区間に対するクエリ（最小値、最大値、総和などの取得）」の双方を、それぞれ $O(\log N)$ の高速な計算量で行うことができる高度なデータ構造です。

## 主なユースケース（どういう時に使うか）

- 「範囲 $[l, r)$ の値すべてに $X$ を足す」かつ「範囲 $[l, r)$ の最小値を求める」といったクエリが大量に発生するゲーム、グラフ、描画処理
- 区間更新・区間取得が混在する高度な値管理

## なぜ使うのか（メリット）

- 値の更新を必要になるまで保留（遅延評価）する機構により、愚直にやると $O(N)$ かかる区間一括操作と区間結果取得の両立を、 $O(\log N)$ に抑え込むことができます。

---

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