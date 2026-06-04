package mylib

func CountDivisibleSubstrings(s string, p int) int {
	N := len(s)
	if p == 2 || p == 5 {
		ans := 0
		for i := 0; i < N; i++ {
			val := int(s[i] - '0')
			if val%p == 0 {
				ans += i + 1
			}
		}
		return ans
	}
	m := make(map[int]int)
	cur := 0
	m[cur] = 1
	digit := 1
	ans := 0
	for i := N - 1; i >= 0; i-- {
		cur += digit * int(s[i]-'0')
		cur %= p
		ans += m[cur]
		m[cur]++
		digit = digit * 10 % p
	}
	return ans
}
