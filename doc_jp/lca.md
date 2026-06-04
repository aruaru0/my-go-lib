# LCA — 最小共通祖先 (Binary Lifting)

## 型

```go
type LCA struct { ... }
```

## 関数

```go
func NewLCA(n int) *LCA
func (l *LCA) AddEdge(from, to int)
func (l *LCA) Build(root int)
func (l *LCA) LCA(u, v int) int
func (l *LCA) Dist(u, v int) int
```

LCA は、Binary Lifting (ダブリング) を使用して根付き木上の最小共通祖先クエリを提供します。

- NewLCA は n 個のノードからなる木を作成します。
- AddEdge は無向辺を追加します。
- Build は祖先テーブルを事前計算します (すべての辺を追加した後、クエリの前に呼び出す必要があります)。
- LCA は u と v の LCA を返します。
- Dist は u と v の間の辺の数を返します。
