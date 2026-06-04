# OrderedMap — Map with Insertion Order

Maintains key-value pairs in insertion order using a slice + map combination.

## Functions

- `NewOrderedMap[K comparable, V any]() *OrderedMap[K, V]` — creates an empty ordered map
- `(*OrderedMap[K, V]).Set(key K, value V)` — sets a key (preserving original insertion order)
- `(*OrderedMap[K, V]).Get(key K) (V, bool)` — gets a value by key
- `(*OrderedMap[K, V]).Delete(key K)` — removes a key
- `(*OrderedMap[K, V]).Keys() []K` — returns keys in insertion order
- `(*OrderedMap[K, V]).Len() int` — returns the number of entries
