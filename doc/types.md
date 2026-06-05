# Types — Generics Type Constraints

Defines generic type constraints used throughout the library.

## Constraints

- `Ordered` — Types that support ordered comparisons (`<`, `>`, `<=`, `>=`). Includes integers, unsigned integers, floats, and strings.
- `Number` — Types that support numeric arithmetic. Includes integers, unsigned integers, and floats.

## Definition

```go
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}
```
