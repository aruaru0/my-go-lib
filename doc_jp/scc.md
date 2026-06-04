# SccGraph (強連結成分分解)

`SccGraph` は、Tarjan のアルゴリズムを使用して有向グラフの強連結成分を計算します。

## コンストラクタ

```go
func NewSccGraph(n int) *SccGraph
```

`n` 個の頂点 (0..n-1) を持つ新しいグラフを作成します。

## メソッド

```go
func (scc *SccGraph) AddEdge(from, to int)
```
`from` から `to` への有向辺を追加します。

```go
func (scc *SccGraph) Scc() [][]int
```
強連結成分のリストを返します。各成分は頂点インデックスのスライスです。成分はトポロジカル順序 (逆 DAG 順) で返されます。

## 使用例

```go
g := NewSccGraph(4)
g.AddEdge(0, 1)
g.AddEdge(1, 2)
g.AddEdge(2, 0)
g.AddEdge(1, 3)
scc := g.Scc()
fmt.Println(scc) // [[3] [0 1 2]]
```

## 計算量

O(V + E) 時間、O(V + E) 空間です。
