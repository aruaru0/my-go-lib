# ZAlgorithm — Z-Algorithm (Generic)

## Functions

- `ZAlgorithm[T comparable](s []T) []int` — computes the Z-array of a comparable slice
- `ZAlgorithmString(s string) []int` — computes the Z-array of a string

The Z-array `z` where `z[i]` is the length of the longest common prefix of `s` and `s[i:]`.

## Example

```go
z := mylib.ZAlgorithmString("aaaaa")
// z = [5 4 3 2 1]

z2 := mylib.ZAlgorithmString("abcabcabc")
// z2 = [9 0 0 6 0 0 3 0 0]
```
