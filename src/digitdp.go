package mylib

import "strconv"

func CountContaining(n int, target int) int64 {
	if n <= 0 {
		return 0
	}
	s := strconv.Itoa(n)
	m := len(s)
	var dp [20][2][2]int64
	dp[0][1][0] = 1
	for i := 0; i < m; i++ {
		d := int(s[i] - '0')
		for tight := 0; tight < 2; tight++ {
			for has := 0; has < 2; has++ {
				cur := dp[i][tight][has]
				if cur == 0 {
					continue
				}
				limit := 9
				if tight == 1 {
					limit = d
				}
				for dig := 0; dig <= limit; dig++ {
					nt := 0
					if tight == 1 && dig == limit {
						nt = 1
					}
					nh := has
					if dig == target {
						nh = 1
					}
					dp[i+1][nt][nh] += cur
				}
			}
		}
	}
	return dp[m][0][1] + dp[m][1][1]
}

func CountNonZero(s string, k int) int64 {
	n := len(s)
	var dp [101][2][5]int64
	dp[0][0][0] = 1
	for i := 1; i <= n; i++ {
		x := int(s[i-1] - '0')
		for j := 0; j <= k; j++ {
			for d := 0; d < 10; d++ {
				if d == 0 {
					dp[i][1][j] += dp[i-1][1][j]
				} else if j > 0 {
					dp[i][1][j] += dp[i-1][1][j-1]
				}
			}
			for d := 0; d < x; d++ {
				if d == 0 {
					dp[i][1][j] += dp[i-1][0][j]
				} else if j > 0 {
					dp[i][1][j] += dp[i-1][0][j-1]
				}
			}
			if x == 0 {
				dp[i][0][j] = dp[i-1][0][j]
			} else if j > 0 {
				dp[i][0][j] = dp[i-1][0][j-1]
			}
		}
	}
	res := dp[n][0][k] + dp[n][1][k]
	if k == 0 {
		res--
	}
	return res
}
