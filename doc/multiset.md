# MultiSet — Generic Multiset

Simple multiset backed by a `map[T]int`. Supports adding and removing elements, tracking counts.

## Functions

- `NewMultiSet[T comparable]() *MultiSet[T]` — creates an empty multiset
- `(*MultiSet[T]).Put(x T)` — adds one instance of `x`
- `(*MultiSet[T]).Remove(x T)` — removes one instance of `x`
- `(*MultiSet[T]).Count(x T) int` — returns the count of `x`
- `(*MultiSet[T]).Values() []T` — returns all distinct elements
- `(*MultiSet[T]).Len() int` — returns the number of distinct elements
