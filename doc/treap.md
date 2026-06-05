# Treap — Randomized Binary Search Tree

Implements a treap (randomized BST) with support for duplicate counts.

## Functions

- `NewTreap[T any](less func(a, b T) bool) *Treap[T]` — creates a new treap with the given ordering
- `(*Treap[T]).Insert(v T)` — inserts a value (increments count if duplicate)
- `(*Treap[T]).Delete(v T)` — removes one instance of the value
- `(*Treap[T]).Find(v T) bool` — returns whether the value exists
- `(*Treap[T]).Min() T` — returns the minimum value
- `(*Treap[T]).Max() T` — returns the maximum value
- `(*Treap[T]).Kth(k int) T` — returns the k-th smallest distinct element (0-indexed)
- `(*Treap[T]).Len() int` — returns the number of distinct keys
