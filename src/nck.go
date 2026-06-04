package mylib

type NCR struct {
	mod   int
	frac  []int
	ifrac []int
}

func NewNCR(mod int, maxN int) *NCR {
	nc := &NCR{
		mod:   mod,
		frac:  make([]int, maxN+1),
		ifrac: make([]int, maxN+1),
	}
	nc.frac[0] = 1
	for i := 1; i <= maxN; i++ {
		nc.frac[i] = nc.frac[i-1] * i % mod
	}
	nc.ifrac[maxN] = mpow(nc.frac[maxN], mod-2, mod)
	for i := maxN - 1; i >= 0; i-- {
		nc.ifrac[i] = nc.ifrac[i+1] * (i + 1) % mod
	}
	return nc
}

func mpow(p, n, mod int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret = ret * x % mod
		}
		n /= 2
		x = x * x % mod
	}
	return ret
}

func (nc *NCR) N(n, k int) int {
	if n < k || k < 0 {
		return 0
	}
	return nc.frac[n] * nc.ifrac[k] % nc.mod * nc.ifrac[n-k] % nc.mod
}

func (nc *NCR) P(n, k int) int {
	if k < 0 || n < k {
		return 0
	}
	return nc.frac[n] * nc.ifrac[n-k] % nc.mod
}

func (nc *NCR) H(n, k int) int {
	if n == 0 && k == 0 {
		return 1
	}
	return nc.N(n+k-1, k)
}
