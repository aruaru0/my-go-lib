# Dinic

`Dinic` computes the maximum flow in a directed graph using the Dinic algorithm. A lean implementation without edge tracking.

## Constructor

```go
func NewDinic(N int) *Dinic
```

Creates a new flow network with `N` vertices (0..N-1).

## Methods

```go
func (d *Dinic) AddEdge(from, to, cap int)
```
Adds a directed edge with capacity `cap`.

```go
func (d *Dinic) MaxFlow(s, t int) int
```
Computes the maximum flow from `s` to `t`.

## Example

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

## Complexity

O(V²·E) time, O(V + E) space.
