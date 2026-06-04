package mylib

type RollingHash struct {
	base int
	mod  int
	pw   []int
	h    []int
}

func NewRollingHash(s string, base, mod int) *RollingHash {
	ret := &RollingHash{base: base, mod: mod}
	n := len(s) + 1
	ret.pw = make([]int, n)
	ret.h = make([]int, n)
	v := 0
	for i := 0; i < len(s); i++ {
		v = (v*base + int(s[i])) % mod
		ret.h[i+1] = v
	}
	v = 1
	for i := 0; i < len(s); i++ {
		v = v * base % mod
		ret.pw[i+1] = v
	}
	return ret
}

func NewRollingHashDefault(s string) *RollingHash {
	return NewRollingHash(s, 37, int(1e9+7))
}

func (rh *RollingHash) Get(l, r int) int {
	ret := (rh.h[r] - rh.h[l]*rh.pw[r-l]) % rh.mod
	if ret < 0 {
		ret += rh.mod
	}
	return ret
}

func (rh *RollingHash) Len() int {
	return len(rh.h) - 1
}
