# ConvexHullTrick — Li Chao Segment Tree

直線 `y = a*x + b` の集合を管理し、指定された `x` での最大値をクエリします。

## 関数

- `NewConvexHullTrick[T Ordered](pos []T) *ConvexHullTrick[T]` — x 座標 `pos` (内部でソート・重複除去されます) に対する木を作成します
- `(*ConvexHullTrick[T]).AddLine(a, b T)` — 直線 `a*x + b` を追加します
- `(*ConvexHullTrick[T]).GetMax(x T) T` — すべての直線における x での最大値を返します
