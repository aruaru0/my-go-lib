package mylib

type NCRGenerator struct {
	mod       int
	fracMemo  []int
	ifracMemo []int
}

func NewNCRGenerator(mod int) *NCRGenerator {
	return &NCRGenerator{
		mod:       mod,
		fracMemo:  []int{1, 1},
		ifracMemo: []int{1, 1},
	}
}

func (g *NCRGenerator) mpow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		tmp := g.mpow(a, b/2)
		return tmp * tmp % g.mod
	}
	return g.mpow(a, b-1) * a % g.mod
}

func (g *NCRGenerator) mfrac(n int) int {
	if len(g.fracMemo) > n {
		return g.fracMemo[n]
	}
	if len(g.fracMemo) == 0 {
		g.fracMemo = append(g.fracMemo, 1)
	}
	for len(g.fracMemo) <= n {
		size := len(g.fracMemo)
		g.fracMemo = append(g.fracMemo, g.fracMemo[size-1]*size%g.mod)
	}
	return g.fracMemo[n]
}

func (g *NCRGenerator) mifrac(n int) int {
	if len(g.ifracMemo) > n {
		return g.ifracMemo[n]
	}
	if len(g.ifracMemo) == 0 {
		g.ifracMemo = append(g.ifracMemo, 1)
	}
	for len(g.ifracMemo) <= n {
		size := len(g.ifracMemo)
		g.ifracMemo = append(g.ifracMemo, g.ifracMemo[size-1]*g.mpow(size, g.mod-2)%g.mod)
	}
	return g.ifracMemo[n]
}

func (g *NCRGenerator) NCR(n, r int) int {
	if n == r {
		return 1
	}
	if n < r || r < 0 {
		return 0
	}
	ret := 1
	ret = ret * g.mfrac(n) % g.mod
	ret = ret * g.mifrac(r) % g.mod
	ret = ret * g.mifrac(n-r) % g.mod
	return ret
}
