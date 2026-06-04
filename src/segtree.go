package mylib

type SegTree[T any] struct {
	n      int
	size   int
	log    int
	d      []T
	e      func() T
	merger func(a, b T) T
}

func NewSegTree[T any](v []T, e func() T, m func(a, b T) T) *SegTree[T] {
	seg := new(SegTree[T])
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]T, 2*seg.size)
	seg.e = e
	seg.merger = m
	for i := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = v[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.update(i)
	}
	return seg
}

func (seg *SegTree[T]) update(k int) {
	seg.d[k] = seg.merger(seg.d[2*k], seg.d[2*k+1])
}

func (seg *SegTree[T]) Set(p int, x T) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.update(p >> uint(i))
	}
}

func (seg *SegTree[T]) Get(p int) T {
	return seg.d[p+seg.size]
}

func (seg *SegTree[T]) Prod(l, r int) T {
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.merger(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.merger(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.merger(sml, smr)
}

func (seg *SegTree[T]) AllProd() T {
	return seg.d[1]
}

func (seg *SegTree[T]) MaxRight(l int, cmp func(T) bool) int {
	if l == seg.n {
		return seg.n
	}
	l += seg.size
	sm := seg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(seg.merger(sm, seg.d[l])) {
			for l < seg.size {
				l = 2 * l
				if cmp(seg.merger(sm, seg.d[l])) {
					sm = seg.merger(sm, seg.d[l])
					l++
				}
			}
			return l - seg.size
		}
		sm = seg.merger(sm, seg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return seg.n
}

func (seg *SegTree[T]) MinLeft(r int, cmp func(T) bool) int {
	if r == 0 {
		return 0
	}
	r += seg.size
	sm := seg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(seg.merger(seg.d[r], sm)) {
			for r < seg.size {
				r = 2*r + 1
				if cmp(seg.merger(seg.d[r], sm)) {
					sm = seg.merger(seg.d[r], sm)
					r--
				}
			}
			return r + 1 - seg.size
		}
		sm = seg.merger(seg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (seg *SegTree[T]) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}
