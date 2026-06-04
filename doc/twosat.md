# TwoSat (2-SAT Solver)

`TwoSat` solves the 2-Satisfiability problem. Given Boolean variables and clauses of the form `(x_i = f) ∨ (x_j = g)`, it determines if a satisfying assignment exists and finds one.

## Constructor

```go
func NewTwoSat(n int) *TwoSat
```

Creates a new 2-SAT instance with `n` Boolean variables (0..n-1).

## Methods

```go
func (ts *TwoSat) AddClause(i int, f bool, j int, g bool)
```
Adds a clause `(x_i == f) ∨ (x_j == g)`.

```go
func (ts *TwoSat) Satisfiable() bool
```
Returns `true` if a satisfying assignment exists.

```go
func (ts *TwoSat) Answer() []bool
```
Returns the satisfying assignment (valid only if `Satisfiable()` returned `true`).

## Example

```go
ts := NewTwoSat(2)
ts.AddClause(0, true, 1, true)
ts.AddClause(0, false, 1, false)
if ts.Satisfiable() {
    fmt.Println(ts.Answer()) // [true true]
}
```

## Complexity

O(N + M) time, O(N + M) space where N is variables and M is clauses.
