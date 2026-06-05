# FenwickTree — ジェネリック Fenwick 木 (Binary Indexed Tree)

配列の要素の更新と、任意の区間の和の計算を、それぞれ $O(\log N)$ の高速な計算量で行うことができるデータ構造です（Fenwick Tree / BIT）。

## 主なユースケース（どういう時に使うか）

- 頻繁に要素が書き換わる動的な配列における、区間の総和やプレフィックス和の算出
- 平面走査法や二次元座標上のイベント処理に伴うカウント管理

## なぜ使うのか（メリット）

- セグメント木に比べて実装がコンパクトであり、メモリ効率に優れ、定数倍の処理速度が非常に速いのが最大の特徴です。

---

## 関数

- `NewFenwickTree[T Integer](n int) *FenwickTree[T]` — サイズ n の木を作成します
- `(*FenwickTree[T]).Add(i int, x T)` — インデックス i の要素に x を加算します
- `(*FenwickTree[T]).Sum(i int) T` — [0, i] (両端を含む) の要素の和を返します。負の i の場合は 0 を返します。
- `(*FenwickTree[T]).RangeSum(l, r int) T` — [l, r) の要素の和を返します
- `(*FenwickTree[T]).LowerBound(x T) int` — `Sum(index) >= x` となる最小のインデックスを返します