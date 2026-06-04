package mylib

func EulerPhi[T Integer](n T) T {
	ret := n
	for i := T(2); i*i <= n; i++ {
		if n%i == 0 {
			ret -= ret / i
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		ret -= ret / n
	}
	return ret
}
