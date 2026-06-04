# UnionFind (Disjoint Set Union)

`UnionFind` is a data structure that manages a partition of a set into disjoint subsets. It supports two operations:

- **Unite**: merge two subsets into one
- **Same**: check if two elements belong to the same subset
- **Size**: get the size of the subset an element belongs to

## Constructor

```go
func NewUnionFind(N int) *UnionFind
```

Creates a new UnionFind with `N` elements (0..N-1), each initially in its own singleton set.

## Methods

```go
func (p *UnionFind) Unite(x, y int)
```
Merges the sets containing `x` and `y`. Does nothing if they are already in the same set.

```go
func (p *UnionFind) Same(x, y int) bool
```
Returns `true` if `x` and `y` belong to the same set.

```go
func (p *UnionFind) Size(x int) int
```
Returns the number of elements in the set containing `x`.

## Example

```go
uf := NewUnionFind(10)
uf.Unite(0, 1)
uf.Unite(1, 2)
fmt.Println(uf.Same(0, 2)) // true
fmt.Println(uf.Size(0))    // 3
```

## Complexity

Both `Unite` and `Same` run in nearly O(1) amortized time (inverse Ackermann function) thanks to path compression and union by size.
