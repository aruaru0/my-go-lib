package mylib

type CombTable[T Integer] struct {
	table [][]T
}

func NewCombTable[T Integer](n int) *CombTable[T] {
	table := make([][]T, n+1)
	for i := range table {
		table[i] = make([]T, n+1)
	}
	table[0][0] = T(1)
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				table[i][j] = T(1)
			} else {
				table[i][j] = table[i-1][j-1] + table[i-1][j]
			}
		}
	}
	return &CombTable[T]{table: table}
}

func (c *CombTable[T]) NCk(n, k int) T {
	if n < k || k < 0 || n >= len(c.table) {
		return 0
	}
	return c.table[n][k]
}
