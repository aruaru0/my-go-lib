# SegTree — Generic Segment Tree

Generic segment tree for any type `T` with arbitrary monoid operations.

## Functions

- `NewSegTree[T any](v []T, e func() T, m func(a, b T) T) *SegTree[T]` — creates a tree from initial slice `v` with identity `e()` and merge operation `m`
- `(*SegTree[T]).Set(p int, x T)` — sets element at index `p` to `x`
- `(*SegTree[T]).Get(p int) T` — gets element at index `p`
- `(*SegTree[T]).Prod(l, r int) T` — returns `m`-product over `[l, r)`
- `(*SegTree[T]).AllProd() T` — returns `m`-product over all elements
- `(*SegTree[T]).MaxRight(l int, cmp func(T) bool) int` — maximum index `r` such that `Prod(l, r)` satisfies `cmp`
- `(*SegTree[T]).MinLeft(r int, cmp func(T) bool) int` — minimum index `l` such that `Prod(l, r)` satisfies `cmp`
