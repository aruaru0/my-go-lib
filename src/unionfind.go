package mylib

type UnionFind struct {
	par []int
	siz []int
}

func NewUnionFind(N int) *UnionFind {
	u := &UnionFind{
		par: make([]int, N),
		siz: make([]int, N),
	}
	for i := 0; i < N; i++ {
		u.par[i] = i
		u.siz[i] = 1
	}
	return u
}

func (p *UnionFind) root(x int) int {
	if p.par[x] != x {
		p.par[x] = p.root(p.par[x])
	}
	return p.par[x]
}

func (p *UnionFind) Unite(x, y int) {
	x = p.root(x)
	y = p.root(y)
	if x == y {
		return
	}
	if p.siz[x] < p.siz[y] {
		x, y = y, x
	}
	p.par[y] = x
	p.siz[x] += p.siz[y]
}

func (p *UnionFind) Same(x, y int) bool {
	return p.root(x) == p.root(y)
}

func (p *UnionFind) Size(x int) int {
	return p.siz[p.root(x)]
}
