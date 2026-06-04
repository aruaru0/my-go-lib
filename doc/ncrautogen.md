# ncrautogen

Auto-growing factorial table for nCr computation.

## Types

```go
type NCRGenerator struct { ... }
```

## Functions

```go
func NewNCRGenerator(mod int) *NCRGenerator
```

## Methods

```go
func (g *NCRGenerator) NCR(n, r int) int
```

Computes `nCr` modulo `mod`. Factorial tables grow automatically as needed.

## Example

```go
g := mylib.NewNCRGenerator(1000000007)
c := g.NCR(5, 2) // 10
c = g.NCR(10, 3) // 120
c = g.NCR(100, 50) // auto-grows tables
```
