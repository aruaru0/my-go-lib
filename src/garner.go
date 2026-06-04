package mylib

func Garner(x, m []int, mod int) int {
	if len(x) == 0 {
		return 0
	}
	v := make([]int, len(x))
	v[0] = x[0]
	for i := 1; i < len(x); i++ {
		X := x[i]
		M := 1
		for j := 0; j < i; j++ {
			X -= v[j] * M
			X %= m[i]
			M *= m[j]
			M %= m[i]
		}
		if X < 0 {
			X += m[i]
		}
		v[i] = X * invModForGarner(M, m[i]) % m[i]
	}
	ret := v[0]
	p := 1
	if mod == 0 {
		for i := 1; i < len(x); i++ {
			p *= m[i-1]
			ret += p * v[i]
		}
	} else {
		ret %= mod
		for i := 1; i < len(x); i++ {
			p *= m[i-1]
			p %= mod
			ret += p * v[i]
			ret %= mod
		}
	}
	return ret
}

func invModForGarner(a, m int) int {
	s := a % m
	t := m
	sx, sy, tx, ty := 1, 0, 0, 1
	for s%t != 0 {
		f := s / t
		u := s - t*f
		ux := sx - tx*f
		uy := sy - ty*f
		s = t
		sx = tx
		sy = ty
		t = u
		tx = ux
		ty = uy
	}
	if tx < 0 {
		tx += m
	}
	return tx
}
