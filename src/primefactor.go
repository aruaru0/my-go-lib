package mylib

func PrimeFactorMap[T Integer](n T) map[T]T {
	pfs := make(map[T]T)
	for n%2 == 0 {
		pfs[2]++
		n /= 2
	}
	for i := T(3); i*i <= n; i += 2 {
		for n%i == 0 {
			pfs[i]++
			n /= i
		}
	}
	if n > 2 {
		pfs[n]++
	}
	return pfs
}

func PrimeFactors[T Integer](n T) []T {
	pfs := make([]T, 0)
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n /= 2
	}
	for i := T(3); i*i <= n; i += 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n /= i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}
	return pfs
}
