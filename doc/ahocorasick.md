# AhoCorasick — Multi-Pattern String Matching

## Types

- `AhoCorasick` — Aho-Corasick automaton
- `AhoMatch` — represents a matched pattern at a position

## Functions

- `NewAhoCorasick() *AhoCorasick` — creates a new automaton
- `(*AhoCorasick).Add(pattern string) int` — adds a pattern, returns its ID
- `(*AhoCorasick).Build()` — builds failure links (call after adding all patterns)
- `(*AhoCorasick).Search(text string) []AhoMatch` — returns all matches (pos, id)
- `(*AhoCorasick).MatchCount(text string) int` — returns total number of matches
- `(*AhoCorasick).GetMask(v int) int` — returns the pattern match mask for state `v`

## Example

```go
ac := mylib.NewAhoCorasick()
ac.Add("he")
ac.Add("she")
ac.Add("hers")
ac.Build()
matches := ac.Search("ushers")
// matches: positions 1 (she), 2 (he), 2 (hers)
```
