# MaxFlow

`MaxFlow` は、Dinic アルゴリズムを使用して有向グラフの最大流を計算します。辺追跡と最小カットクエリをサポートします。

## 型

```go
type MFEdge struct {
    From, To, Capa, Flow int
}
```

## コンストラクタ

```go
func NewMaxFlow(n int) *MaxFlow
```

`n` 個の頂点 (0..n-1) を持つ新しいフローネットワークを作成します。

## メソッド

```go
func (mf *MaxFlow) AddEdge(from, to, capa int) int
```
容量 `capa` の有向辺を追加します。辺のインデックスを返します。

```go
func (mf *MaxFlow) Flow(s, t int) int
```
`s` から `t` への最大流を計算します。

```go
func (mf *MaxFlow) FlowL(s, t, flowLim int) int
```
流量制限付きの最大流を計算します。

```go
func (mf *MaxFlow) MinCut(s int) []bool
```
残余グラフにおいて `s` から到達可能な頂点を示すブールスライスを返します (最小 s-t カット)。

```go
func (mf *MaxFlow) GetEdge(i int) MFEdge
```
インデックスで辺の情報を返します。

```go
func (mf *MaxFlow) EdgesList() []MFEdge
```
すべての辺を返します。

```go
func (mf *MaxFlow) ChangeEdge(i, newCapa, newFlow int)
```
辺の容量と流量を変更します。

## 使用例

```go
mf := NewMaxFlow(4)
mf.AddEdge(0, 1, 10)
mf.AddEdge(0, 2, 5)
mf.AddEdge(1, 2, 15)
mf.AddEdge(1, 3, 10)
mf.AddEdge(2, 3, 10)
flow := mf.Flow(0, 3)
fmt.Println(flow) // 15
```

## 計算量

O(V²·E) 時間、O(V + E) 空間です。
