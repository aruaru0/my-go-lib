# digitdp

Digit DP utilities for counting numbers with specific digit properties.

## Functions

```go
func CountContaining(n int, target int) int64
func CountNonZero(s string, k int) int64
```

## CountContaining

Counts numbers from 1 to n (inclusive) that contain the digit `target`.

## CountNonZero

Counts numbers from 1 to s (inclusive, where s is a decimal string) that have
exactly `k` non-zero digits. Based on ABC 154 E.
