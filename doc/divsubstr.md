# divsubstr

Count substrings of a decimal string divisible by P.

## Functions

```go
func CountDivisibleSubstrings(s string, p int) int
```

## CountDivisibleSubstrings

Counts the number of substrings of `s` that are divisible by `p`.
Uses a right-to-left modulo accumulation with a map to track prefix remainders.
For `p = 2` or `p = 5`, uses a special optimized path.
