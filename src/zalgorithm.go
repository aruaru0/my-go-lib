package mylib

func ZAlgorithm[T comparable](s []T) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	z := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		if j+z[j] <= i {
			z[i] = 0
		} else {
			if j+z[j]-i < z[i-j] {
				z[i] = j + z[j] - i
			} else {
				z[i] = z[i-j]
			}
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if j+z[j] < i+z[i] {
			j = i
		}
	}
	z[0] = n
	return z
}

func ZAlgorithmString(s string) []int {
	return ZAlgorithm([]byte(s))
}
