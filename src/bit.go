package mylib

type BIT struct {
	n int
	d []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n, d: make([]int, n+1)}
}

func (b *BIT) Add(k, v int) {
	k++
	for k <= b.n {
		b.d[k] += v
		k += k & -k
	}
}

func (b *BIT) Sum(k int) int {
	if k < 0 {
		return 0
	}
	k++
	if k > b.n {
		k = b.n
	}
	ret := 0
	for k > 0 {
		ret += b.d[k]
		k -= k & -k
	}
	return ret
}

func (b *BIT) RangeSum(l, r int) int {
	if l >= r {
		return 0
	}
	return b.Sum(r-1) - b.Sum(l-1)
}
