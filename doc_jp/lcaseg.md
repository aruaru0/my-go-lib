# LCASeg — 最小共通祖先 (Segment Tree / Euler Tour)

木のオイラーツアー（Euler Tour）を行い、訪問順の頂点深度配列に対する区間最小値クエリ（RMQ）問題に変換することで、セグメント木を用いて最近共通祖先（LCA）を探索するアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- オイラーツアー結果をセグメント木やRMQ用のデータ構造と共有し、LCAを算出したい場合
- 木の頂点更新などが伴い、動的にLCAや距離を管理したい場合

## なぜ使うのか（メリット）

- オイラーツアーにより木を平坦な配列に変換するため、セグメント木等の既存のRMQライブラリをそのまま活用して $O(N)$ の構築および $O(\log N)$ のクエリを実現できます。

---

## 型

```go
type LCASeg struct { ... }
```

## 関数

```go
func NewLCASeg(root, n int, node [][]int) *LCASeg
func (l *LCASeg) LCA(u, v int) int
func (l *LCASeg) Dist(u, v int) int
```

LCASeg は、Euler ツアー + Segment Tree (深さ配列に対する RMQ) を使用して LCA クエリを提供します。

- NewLCASeg は、指定された根、ノード数、隣接リストから構造体を作成します。
- LCA は u と v の LCA を返します。
- Dist は u と v の間の辺の数を返します。