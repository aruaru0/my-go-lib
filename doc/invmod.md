# invmod

Precomputes modular inverses in O(n).

## Functions

```go
func InvModTable(n, m int) []int
```

Returns an array of length `n` where `ret[i] = i^{-1} mod m` for `1 <= i < n`. Uses the recurrence `inv[i] = m - m/i * inv[m%i] % m`.

## Example

```go
inv := mylib.InvModTable(10, 13)
// inv[1..9] are modular inverses modulo 13
```
