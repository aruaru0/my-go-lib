# ConvexHullTrick — Li Chao Segment Tree

複数の直線の式 $y = a_i x + b_i$ が与えられたとき、任意の $x$ に対する $y$ の最小値（または最大値）を高速にクエリ・管理するアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- **動的計画法（DP）の高速化（例: $DP[i] = \min_{j < i} (DP[j] + A[i] \cdot B[j] + C[i])$ のような形の遷移）**
- 投資やエネルギー消費の最適化問題で、時間経過に伴う複数の線形コストモデルの中から常に最小のものを選択する

## なぜ使うのか（メリット）

- 直線を素朴に走査すると $O(N)$ かかるクエリを、直線の傾きや追加順の性質を利用して $O(\log N)$ またはクエリ単調性がある場合はならし $O(1)$ で極めて高速に計算できます。

---

## 関数

- `NewConvexHullTrick[T Ordered](pos []T) *ConvexHullTrick[T]` — x 座標 `pos` (内部でソート・重複除去されます) に対する木を作成します
- `(*ConvexHullTrick[T]).AddLine(a, b T)` — 直線 `a*x + b` を追加します
- `(*ConvexHullTrick[T]).GetMax(x T) T` — すべての直線における x での最大値を返します