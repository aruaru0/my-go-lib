package mylib

func Mul3x3(A, B [3][3]int) [3][3]int {
	var C [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func MulMatrix[T Integer](A, B [][]T) [][]T {
	n, m := len(A), len(A[0])
	p := len(B[0])
	C := make([][]T, n)
	for i := 0; i < n; i++ {
		C[i] = make([]T, p)
		for j := 0; j < p; j++ {
			var sum T
			for k := 0; k < m; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C
}
