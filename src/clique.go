package mylib

const cliqueInf = 255

func MinCliqueCover(N int, edges [][]int) int {
	G := make([]int, N)
	for i := 0; i < N; i++ {
		for _, j := range edges[i] {
			G[i] |= 1 << j
		}
	}

	n := 1 << N
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = cliqueInf
	}
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < N; j++ {
			if i&(1<<j) == 0 {
				continue
			}
			t := i ^ (1 << j)
			if dp[t] == 1 && G[j]&t == t {
				dp[i] = 1
			}
			break
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j != 0; j = (j - 1) & i {
			dp[i] = min(dp[i], dp[j]+dp[i^j])
		}
	}
	return dp[n-1]
}
