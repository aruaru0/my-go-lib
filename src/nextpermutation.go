package mylib

func NextPermutation[T Ordered](x []T) bool {
	n := len(x) - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; j >= 0 && !(x[j] < x[j+1]); j-- {
	}
	if j < 0 {
		return false
	}
	l := n
	for !(x[j] < x[l]) {
		l--
	}
	x[j], x[l] = x[l], x[j]
	for k, l := j+1, n; k < l; {
		x[k], x[l] = x[l], x[k]
		k++
		l--
	}
	return true
}
