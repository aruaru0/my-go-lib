# MP — String Pattern Matching (Morris-Pratt Algorithm)

## Types

- `MP` — pattern matcher struct

## Functions

- `NewMP(pattern string) *MP` — creates a new pattern matcher
- `(*MP).FindAll(text string) []int` — returns all start positions of the pattern in text

## Example

```go
mp := mylib.NewMP("abc")
pos := mp.FindAll("abcabcabc")
// pos = [0, 3, 6]
```
