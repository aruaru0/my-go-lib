package mylib

func Divisors[T Integer](n T) []T {
	var ret []T
	for i := T(1); i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i*i != n {
				ret = append(ret, n/i)
			}
		}
	}
	return ret
}
