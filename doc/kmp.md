# KMP — Generic Knuth-Morris-Pratt Pattern Matching

## Types

- `KMP[T comparable]` — generic pattern matcher for any comparable slice type

## Functions

- `NewKMP[T comparable](pattern []T) *KMP[T]` — creates a new pattern matcher
- `(*KMP[T]).Search(text []T) []int` — returns all start positions of the pattern in text

## Example

```go
kmp := mylib.NewKMP([]int{1, 2, 3})
pos := kmp.Search([]int{1, 2, 3, 1, 2, 3})
// pos = [0, 3]

kmpStr := mylib.NewKMP([]byte("abc"))
posStr := kmpStr.Search([]byte("abcabc"))
// posStr = [0, 3]
```
