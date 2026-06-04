package mylib

type Modint struct {
	mod       int
	factMemo  []int
	ifactMemo []int
}

func NewModint(m int) *Modint {
	return &Modint{
		mod:       m,
		factMemo:  []int{1, 1},
		ifactMemo: []int{1, 1},
	}
}

func (m *Modint) Add(a, b int) int {
	ret := (a + b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *Modint) Sub(a, b int) int {
	ret := (a - b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *Modint) Mul(a, b int) int {
	ret := a * b % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *Modint) Div(a, b int) int {
	return m.Mul(a, m.Inv(b))
}

func (m *Modint) Pow(p, n int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret = ret * x % m.mod
		}
		n /= 2
		x = x * x % m.mod
	}
	return ret
}

func (m *Modint) Inv(a int) int {
	b, u, v := m.mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m.mod
	if u < 0 {
		u += m.mod
	}
	return u
}

func (m *Modint) MulMatrix(A, B [][]int) [][]int {
	H := len(A)
	W := len(B[0])
	K := len(A[0])
	C := make([][]int, H)
	for i := 0; i < H; i++ {
		C[i] = make([]int, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
				C[i][j] %= m.mod
			}
		}
	}
	return C
}

func (m *Modint) PowMatrix(A [][]int, p int) [][]int {
	N := len(A)
	ret := make([][]int, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]int, N)
		ret[i][i] = 1
	}
	for p > 0 {
		if p&1 == 1 {
			ret = m.MulMatrix(ret, A)
		}
		A = m.MulMatrix(A, A)
		p >>= 1
	}
	return ret
}
