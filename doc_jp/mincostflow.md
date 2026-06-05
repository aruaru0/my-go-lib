# MinCostFlow

各辺に容量だけでなく「単位流量あたりのコスト」が設定されたネットワークにおいて、始点から終点へ指定された流量 $F$ を流すために必要な最小の合計コストを求めるアルゴリズムです（Primal-Dual法）。

## 主なユースケース（どういう時に使うか）

- 物流ネットワークにおいて、輸送費用を最小限に抑えつつ必要な物資を配送する最適化
- **重み付き二部マッチング問題（能力や好みが異なるペアの総満足度の最大化・コスト最小化）**

## なぜ使うのか（メリット）

- ポテンシャルを用いたダイクストラ法の繰り返しにより、負の辺がない（あるいは前処理済みの）ネットワークで、流量 $F$ に対して $O(F \cdot E \log V)$ で最小費用流を効率的に求められます。

---

## 型

```go
type MCFEdge struct {
    From, To, Capa, Flow, Cost int
}
```

## コンストラクタ

```go
func NewMinCostFlow(n int) *MinCostFlow
```

`n` 個の頂点 (0..n-1) を持つ新しいフローネットワークを作成します。

## メソッド

```go
func (mcf *MinCostFlow) AddEdge(from, to, capa, cost int) int
```
容量 `capa`、単位流量あたりのコスト `cost` の有向辺を追加します。

```go
func (mcf *MinCostFlow) GetEdge(i int) MCFEdge
```
インデックス `i` の辺の現在の状態を返します。

```go
func (mcf *MinCostFlow) Edges() []MCFEdge
```
すべての辺の現在の状態を返します。

```go
func (mcf *MinCostFlow) Flow(s, t int) [2]int
```
`s` から `t` への最小費用流を計算します。`[flow, cost]` を返します。

```go
func (mcf *MinCostFlow) FlowL(s, t, flowLim int) [2]int
```
流量制限付きの最小費用流を計算します。

```go
func (mcf *MinCostFlow) Slope(s, t int) [][2]int
```
傾き (区分的線形コスト関数) を返します。

```go
func (mcf *MinCostFlow) SlopeL(s, t, flowLim int) [][2]int
```
流量制限付きの傾きを返します。

## 使用例

```go
mcf := NewMinCostFlow(4)
mcf.AddEdge(0, 1, 10, 2)
mcf.AddEdge(0, 2, 10, 5)
mcf.AddEdge(1, 2, 10, 1)
mcf.AddEdge(1, 3, 10, 3)
mcf.AddEdge(2, 3, 10, 1)
flow, cost := mcf.Flow(0, 3)[0], mcf.Flow(0, 3)[1]
fmt.Println(flow, cost)
```

## 計算量

O(F·E log V) です。F は流量です。