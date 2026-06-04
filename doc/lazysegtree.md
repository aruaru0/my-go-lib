# LazySegTree — Generic Lazy Segment Tree

Segment tree with lazy propagation. Supports range updates and range queries.

## Functions

- `NewLazySegTree[S, F any](v []S, e func() S, merger func(a, b S) S, mapper func(f F, x S) S, comp func(f, g F) F, id func() F) *LazySegTree[S, F]` — creates a tree
  - `e`: identity for S
  - `merger`: combine two S values
  - `mapper`: apply an F to an S
  - `comp`: compose two F values
  - `id`: identity for F
- `(*LazySegTree[S, F]).Set(p int, x S)` — point set
- `(*LazySegTree[S, F]).Get(p int) S` — point get
- `(*LazySegTree[S, F]).Prod(l, r int) S` — range query over `[l, r)`
- `(*LazySegTree[S, F]).AllProd() S` — query over all elements
- `(*LazySegTree[S, F]).Apply(p int, f F)` — point apply
- `(*LazySegTree[S, F]).RangeApply(l, r int, f F)` — range apply over `[l, r)`
- `(*LazySegTree[S, F]).MaxRight(l int, cmp func(S) bool) int` — binary search to the right
- `(*LazySegTree[S, F]).MinLeft(r int, cmp func(S) bool) int` — binary search to the left
