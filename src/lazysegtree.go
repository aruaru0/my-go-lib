package mylib

type LazySegTree[S, F any] struct {
	n      int
	size   int
	log    int
	d      []S
	lz     []F
	e      func() S
	merger func(a, b S) S
	mapper func(f F, x S) S
	comp   func(f, g F) F
	id     func() F
}

func NewLazySegTree[S, F any](v []S, e func() S, merger func(a, b S) S, mapper func(f F, x S) S, comp func(f, g F) F, id func() F) *LazySegTree[S, F] {
	lseg := new(LazySegTree[S, F])
	lseg.n = len(v)
	lseg.log = lseg.ceilPow2(lseg.n)
	lseg.size = 1 << uint(lseg.log)
	lseg.d = make([]S, 2*lseg.size)
	lseg.lz = make([]F, lseg.size)
	lseg.e = e
	lseg.merger = merger
	lseg.mapper = mapper
	lseg.comp = comp
	lseg.id = id
	for i := range lseg.d {
		lseg.d[i] = lseg.e()
	}
	for i := range lseg.lz {
		lseg.lz[i] = lseg.id()
	}
	for i := 0; i < lseg.n; i++ {
		lseg.d[lseg.size+i] = v[i]
	}
	for i := lseg.size - 1; i >= 1; i-- {
		lseg.update(i)
	}
	return lseg
}

func (lseg *LazySegTree[S, F]) update(k int) {
	lseg.d[k] = lseg.merger(lseg.d[2*k], lseg.d[2*k+1])
}

func (lseg *LazySegTree[S, F]) allApply(k int, f F) {
	lseg.d[k] = lseg.mapper(f, lseg.d[k])
	if k < lseg.size {
		lseg.lz[k] = lseg.comp(f, lseg.lz[k])
	}
}

func (lseg *LazySegTree[S, F]) push(k int) {
	lseg.allApply(2*k, lseg.lz[k])
	lseg.allApply(2*k+1, lseg.lz[k])
	lseg.lz[k] = lseg.id()
}

func (lseg *LazySegTree[S, F]) Set(p int, x S) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> uint(i))
	}
	lseg.d[p] = x
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> uint(i))
	}
}

func (lseg *LazySegTree[S, F]) Get(p int) S {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> uint(i))
	}
	return lseg.d[p]
}

func (lseg *LazySegTree[S, F]) Prod(l, r int) S {
	if l == r {
		return lseg.e()
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>uint(i))<<uint(i) != l {
			lseg.push(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg.push(r >> uint(i))
		}
	}
	sml, smr := lseg.e(), lseg.e()
	for l < r {
		if (l & 1) == 1 {
			sml = lseg.merger(sml, lseg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = lseg.merger(lseg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return lseg.merger(sml, smr)
}

func (lseg *LazySegTree[S, F]) AllProd() S {
	return lseg.d[1]
}

func (lseg *LazySegTree[S, F]) Apply(p int, f F) {
	p += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(p >> uint(i))
	}
	lseg.d[p] = lseg.mapper(f, lseg.d[p])
	for i := 1; i <= lseg.log; i++ {
		lseg.update(p >> uint(i))
	}
}

func (lseg *LazySegTree[S, F]) RangeApply(l, r int, f F) {
	if l == r {
		return
	}
	l += lseg.size
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		if (l>>uint(i))<<uint(i) != l {
			lseg.push(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg.push((r - 1) >> uint(i))
		}
	}
	l2, r2 := l, r
	for l < r {
		if l&1 == 1 {
			lseg.allApply(l, f)
			l++
		}
		if r&1 == 1 {
			r--
			lseg.allApply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2
	for i := 1; i <= lseg.log; i++ {
		if (l>>uint(i))<<uint(i) != l {
			lseg.update(l >> uint(i))
		}
		if (r>>uint(i))<<uint(i) != r {
			lseg.update((r - 1) >> uint(i))
		}
	}
}

func (lseg *LazySegTree[S, F]) MaxRight(l int, cmp func(S) bool) int {
	if l == lseg.n {
		return lseg.n
	}
	l += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push(l >> uint(i))
	}
	sm := lseg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(lseg.merger(sm, lseg.d[l])) {
			for l < lseg.size {
				lseg.push(l)
				l = 2 * l
				if cmp(lseg.merger(sm, lseg.d[l])) {
					sm = lseg.merger(sm, lseg.d[l])
					l++
				}
			}
			return l - lseg.size
		}
		sm = lseg.merger(sm, lseg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return lseg.n
}

func (lseg *LazySegTree[S, F]) MinLeft(r int, cmp func(S) bool) int {
	if r == 0 {
		return 0
	}
	r += lseg.size
	for i := lseg.log; i >= 1; i-- {
		lseg.push((r - 1) >> uint(i))
	}
	sm := lseg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(lseg.merger(lseg.d[r], sm)) {
			for r < lseg.size {
				lseg.push(r)
				r = 2*r + 1
				if cmp(lseg.merger(lseg.d[r], sm)) {
					sm = lseg.merger(lseg.d[r], sm)
					r--
				}
			}
			return r + 1 - lseg.size
		}
		sm = lseg.merger(lseg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (lseg *LazySegTree[S, F]) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}
