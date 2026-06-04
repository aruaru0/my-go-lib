# SuffixArray — SA-IS Suffix Array Construction

## Functions

- `SuffixArray(s []int, upper int) []int` — suffix array of an integer slice with values in [0, upper]
- `SuffixArrayString(s string) []int` — suffix array of a string (byte values)
- `SuffixArrayInt(s []int) []int` — suffix array of an arbitrary integer slice (internally compressed)
- `LcpArray(s, sa []int) []int` — LCP array from integer slice and its suffix array
- `LcpArrayString(s string, sa []int) []int` — LCP array from string and its suffix array

## Example

```go
sa := mylib.SuffixArrayString("banana")
// sa = [5 3 1 0 4 2]

lcp := mylib.LcpArrayString("banana", sa)
// lcp = [1 3 0 0 2]
```
