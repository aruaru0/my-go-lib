# DSU (Disjoint Set Union)

`DSU` is a data structure that tracks a partition of a set into disjoint subsets, also known as Union-Find. It supports path compression and union by size.

## Constructor

```go
func NewDsu(n int) *DSU
```

Creates a new DSU with `n` elements (0..n-1), each initially in its own singleton set.

## Methods

```go
func (d *DSU) Merge(a, b int) int
```
Merges the sets containing `a` and `b`. Returns the new leader.

```go
func (d *DSU) Same(a, b int) bool
```
Returns `true` if `a` and `b` belong to the same set.

```go
func (d *DSU) Leader(a int) int
```
Returns the representative element of the set containing `a`.

```go
func (d *DSU) Size(a int) int
```
Returns the number of elements in the set containing `a`.

```go
func (d *DSU) Groups() [][]int
```
Returns a list of all groups (connected components).

## Example

```go
d := NewDsu(6)
d.Merge(0, 1)
d.Merge(2, 3)
d.Merge(0, 2)
fmt.Println(d.Same(1, 3)) // true
fmt.Println(d.Size(0))    // 4
fmt.Println(d.Groups())   // [[0 1 2 3] [4] [5]]
```

## Complexity

Both `Merge` and `Same` run in nearly O(1) amortized time (inverse Ackermann function).
