# compress

Coordinate compression for ordered types.

## Functions

```go
func Compress[T Ordered](s []T) []int
```

## Compress

Returns an integer slice where each element is the rank (0-based) of the corresponding element in `s`, based on sorted unique values. Supports any type satisfying the `Ordered` constraint (int variants, float variants, string).
