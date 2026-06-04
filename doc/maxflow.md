# MaxFlow

`MaxFlow` computes the maximum flow in a directed graph using the Dinic algorithm. Supports edge tracking and min-cut queries.

## Types

```go
type MFEdge struct {
    From, To, Capa, Flow int
}
```

## Constructor

```go
func NewMaxFlow(n int) *MaxFlow
```

Creates a new flow network with `n` vertices (0..n-1).

## Methods

```go
func (mf *MaxFlow) AddEdge(from, to, capa int) int
```
Adds a directed edge with capacity `capa`. Returns the edge index.

```go
func (mf *MaxFlow) Flow(s, t int) int
```
Computes max flow from `s` to `t`.

```go
func (mf *MaxFlow) FlowL(s, t, flowLim int) int
```
Computes max flow with a flow limit.

```go
func (mf *MaxFlow) MinCut(s int) []bool
```
Returns a boolean slice indicating which vertices are reachable from `s` in the residual graph (minimum s-t cut).

```go
func (mf *MaxFlow) GetEdge(i int) MFEdge
```
Returns edge info by index.

```go
func (mf *MaxFlow) EdgesList() []MFEdge
```
Returns all edges.

```go
func (mf *MaxFlow) ChangeEdge(i, newCapa, newFlow int)
```
Modifies an edge's capacity and flow.

## Example

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

## Complexity

O(V²·E) time, O(V + E) space.
