# LCASeg — 最小共通祖先 (Segment Tree / Euler Tour)

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
