package mylib

func InvModTable(n, m int) []int {
	ret := make([]int, n)
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret[i] = m / i * (-ret[m%i])
		ret[i] %= m
		if ret[i] < 0 {
			ret[i] += m
		}
	}
	return ret
}
