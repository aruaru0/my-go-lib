package mylib

const doublingBits = 60

func KthStep(N int, K int64, next []int) int {
	d := make([][]int, doublingBits)
	for i := 0; i < doublingBits; i++ {
		d[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		d[0][i] = next[i]
	}
	for i := 0; i < doublingBits-1; i++ {
		for j := 0; j < N; j++ {
			d[i+1][j] = d[i][d[i][j]]
		}
	}
	pos := 0
	for i := 0; i < doublingBits; i++ {
		if (K>>i)&1 == 1 {
			pos = d[i][pos]
		}
	}
	return pos
}
