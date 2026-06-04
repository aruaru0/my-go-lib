package mylib

type lcasegData struct {
	val, idx int
}

type lcasegTree struct {
	n      int
	size   int
	log    int
	d      []lcasegData
	e      func() lcasegData
	merger func(a, b lcasegData) lcasegData
}

func newLCASegTree(v []lcasegData, e func() lcasegData, m func(a, b lcasegData) lcasegData) *lcasegTree {
	seg := &lcasegTree{}
	seg.n = len(v)
	seg.log = ceilPow2Len(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]lcasegData, 2*seg.size)
	seg.e = e
	seg.merger = m
	for i := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = v[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.d[i] = seg.merger(seg.d[2*i], seg.d[2*i+1])
	}
	return seg
}

func ceilPow2Len(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

func (seg *lcasegTree) prod(l, r int) lcasegData {
	sml := seg.e()
	smr := seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			sml = seg.merger(sml, seg.d[l])
			l++
		}
		if r&1 == 1 {
			r--
			smr = seg.merger(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.merger(sml, smr)
}

type LCASeg struct {
	G     [][]int
	vs    []int
	depth []int
	id    []int
	n     int
	k     int
	seg   *lcasegTree
}

func NewLCASeg(root, n int, node [][]int) *LCASeg {
	l := &LCASeg{}
	l.n = n
	l.G = node
	l.vs = make([]int, n*2-1)
	l.depth = make([]int, n*2-1)
	l.id = make([]int, n)
	l.k = 0
	l.dfs(root, -1, 0)

	initData := make([]lcasegData, n*2-1)
	for i := range l.depth {
		initData[i] = lcasegData{val: l.depth[i], idx: i}
	}
	e := func() lcasegData {
		return lcasegData{val: 1e16, idx: -1}
	}
	merger := func(a, b lcasegData) lcasegData {
		if a.val < b.val {
			return a
		}
		return b
	}
	l.seg = newLCASegTree(initData, e, merger)
	return l
}

func (l *LCASeg) dfs(v, p, d int) {
	l.id[v] = l.k
	l.vs[l.k] = v
	l.depth[l.k] = d
	l.k++
	for i := 0; i < len(l.G[v]); i++ {
		if l.G[v][i] != p {
			l.dfs(l.G[v][i], v, d+1)
			l.vs[l.k] = v
			l.depth[l.k] = d
			l.k++
		}
	}
}

func (l *LCASeg) LCA(u, v int) int {
	a, b := l.id[u], l.id[v]
	if a > b {
		a, b = b, a
	}
	res := l.seg.prod(a, b+1)
	return l.vs[res.idx]
}

func (l *LCASeg) Dist(u, v int) int {
	lc := l.LCA(u, v)
	return l.depth[l.id[u]] + l.depth[l.id[v]] - 2*l.depth[l.id[lc]]
}
