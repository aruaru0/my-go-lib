# MinCostFlow

`MinCostFlow` computes the minimum cost flow in a directed graph using successive shortest paths (Dijkstra with potentials).

## Types

```go
type MCFEdge struct {
    From, To, Capa, Flow, Cost int
}
```

## Constructor

```go
func NewMinCostFlow(n int) *MinCostFlow
```

Creates a new flow network with `n` vertices (0..n-1).

## Methods

```go
func (mcf *MinCostFlow) AddEdge(from, to, capa, cost int) int
```
Adds a directed edge with capacity `capa` and cost per unit of flow.

```go
func (mcf *MinCostFlow) GetEdge(i int) MCFEdge
```
Returns the current state of the edge at index `i`.

```go
func (mcf *MinCostFlow) Edges() []MCFEdge
```
Returns the current state of all edges.

```go
func (mcf *MinCostFlow) Flow(s, t int) [2]int
```
Computes min-cost flow from `s` to `t`. Returns `[flow, cost]`.

```go
func (mcf *MinCostFlow) FlowL(s, t, flowLim int) [2]int
```
Computes min-cost flow with a flow limit.

```go
func (mcf *MinCostFlow) Slope(s, t int) [][2]int
```
Returns the slope (piecewise linear cost function).

```go
func (mcf *MinCostFlow) SlopeL(s, t, flowLim int) [][2]int
```
Returns the slope with a flow limit.

## Example

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

## Complexity

O(F·E log V) where F is the flow amount.
