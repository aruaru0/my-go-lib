# Set — Generic Set

Type-safe generic set backed by a Go map.

## Functions

- `NewSet[T comparable]() Set[T]` — creates an empty set
- `NewSetFromSlice[T comparable](s []T) Set[T]` — creates a set from a slice
- `(Set[T]).Add(v T) bool` — adds value, returns true if newly added
- `(Set[T]).Remove(v T)` — removes a value
- `(Set[T]).Contains(v T) bool` — checks membership
- `(Set[T]).Cardinality() int` — number of elements
- `(Set[T]).Clear()` — removes all elements
- `(Set[T]).Values() []T` — returns all elements as a slice
- `(Set[T]).Union(other Set[T]) Set[T]` — returns a new union set
- `(Set[T]).Intersect(other Set[T]) Set[T]` — returns a new intersection set
- `(Set[T]).Difference(other Set[T]) Set[T]` — returns a new difference set
- `(Set[T]).Equal(other Set[T]) bool` — checks set equality
