package mylib

func PowMatrix[T Integer](A [][]T, p int, mod T) [][]T {
	N := len(A)
	ret := make([][]T, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]T, N)
		ret[i][i] = 1
	}
	for p > 0 {
		if p&1 == 1 {
			ret = mulMatrixGeneric(ret, A, mod)
		}
		A = mulMatrixGeneric(A, A, mod)
		p >>= 1
	}
	return ret
}

func mulMatrixGeneric[T Integer](A, B [][]T, mod T) [][]T {
	H := len(A)
	W := len(B[0])
	K := len(A[0])
	C := make([][]T, H)
	for i := 0; i < H; i++ {
		C[i] = make([]T, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var sum T
			for k := 0; k < K; k++ {
				sum += A[i][k] * B[k][j]
				sum %= mod
			}
			C[i][j] = sum
		}
	}
	return C
}
