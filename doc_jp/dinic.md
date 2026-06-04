# Dinic

`Dinic` は、Dinic アルゴリズムを使用して有向グラフの最大流を計算します。辺追跡機能のないシンプルな実装です。

## コンストラクタ

```go
func NewDinic(N int) *Dinic
```

`N` 個の頂点 (0..N-1) を持つ新しいフローネットワークを作成します。

## メソッド

```go
func (d *Dinic) AddEdge(from, to, cap int)
```
容量 `cap` の有向辺を追加します。

```go
func (d *Dinic) MaxFlow(s, t int) int
```
`s` から `t` への最大流を計算します。

## 使用例

```go
d := NewDinic(4)
d.AddEdge(0, 1, 10)
d.AddEdge(0, 2, 5)
d.AddEdge(1, 2, 15)
d.AddEdge(1, 3, 10)
d.AddEdge(2, 3, 10)
flow := d.MaxFlow(0, 3)
fmt.Println(flow) // 15
```

## 計算量

O(V²·E) 時間、O(V + E) 空間です。
