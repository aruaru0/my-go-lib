package mylib

type FenwickTree[T Integer] struct {
	v []T
}

func NewFenwickTree[T Integer](n int) *FenwickTree[T] {
	return &FenwickTree[T]{v: make([]T, n)}
}

func (f *FenwickTree[T]) Add(i int, x T) {
	for i++; i <= len(f.v); i += i & -i {
		f.v[i-1] += x
	}
}

func (f *FenwickTree[T]) Sum(i int) T {
	var ret T
	if i < 0 {
		return ret
	}
	for i++; i > 0; i -= i & -i {
		ret += f.v[i-1]
	}
	return ret
}

func (f *FenwickTree[T]) RangeSum(l, r int) T {
	var zero T
	if l >= r {
		return zero
	}
	return f.Sum(r-1) - f.Sum(l-1)
}

func (f *FenwickTree[T]) LowerBound(x T) int {
	idx := 0
	n := len(f.v)
	k := 1
	for k < n {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k <= n && f.v[idx+k-1] < x {
			x -= f.v[idx+k-1]
			idx += k
		}
	}
	return idx
}
