package mylib

import "math/bits"

func Convolution(a, b []int, mod int) []int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 {
		return []int{}
	}
	if min(n, m) <= 60 {
		return convolutionNaive(a, b, mod)
	}
	z := 1 << CeilPow2(n + m - 1)
	aa, bb := make([]int, z), make([]int, z)
	copy(aa, a)
	copy(bb, b)

	butterfly(aa, mod)
	butterfly(bb, mod)
	for i := 0; i < z; i++ {
		aa[i] = aa[i] * bb[i] % mod
	}
	butterflyInv(aa, mod)
	aa = aa[:n+m-1]
	iz := InvGcd(z, mod)[1]
	for i := 0; i < n+m-1; i++ {
		aa[i] = aa[i] * iz % mod
	}
	return aa
}

func ConvolutionLL(a, b []int) []int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 {
		return []int{}
	}
	const (
		MOD1 = 754974721
		MOD2 = 167772161
		MOD3 = 469762049
	)
	M2M3 := uint64(MOD2) * uint64(MOD3)
	M1M3 := uint64(MOD1) * uint64(MOD3)
	M1M2 := uint64(MOD1) * uint64(MOD2)
	M1M2M3 := M1M2 * uint64(MOD3)

	i1 := InvGcd(MOD2*MOD3, MOD1)[1]
	i2 := InvGcd(MOD1*MOD3, MOD2)[1]
	i3 := InvGcd(MOD1*MOD2, MOD3)[1]

	c1 := Convolution(a, b, MOD1)
	c2 := Convolution(a, b, MOD2)
	c3 := Convolution(a, b, MOD3)

	c := make([]int, n+m-1)
	for i := 0; i < n+m-1; i++ {
		x := uint64(0)
		x += uint64((c1[i] * i1) % MOD1) * M2M3
		x += uint64((c2[i] * i2) % MOD2) * M1M3
		x += uint64((c3[i] * i3) % MOD3) * M1M2
		t := int(x % MOD1)
		diff := c1[i] - t
		if diff < 0 {
			diff += MOD1
		}
		offset := []uint64{0, 0, M1M2M3, 2 * M1M2M3, 3 * M1M2M3}
		x -= offset[diff%5]
		c[i] = int(x)
	}
	return c
}

func convolutionNaive(a, b []int, mod int) []int {
	n, m := len(a), len(b)
	if n < m {
		n, m = m, n
		a, b = b, a
	}
	ans := make([]int, n+m-1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans[i+j] += a[i] * b[j] % mod
			ans[i+j] %= mod
		}
	}
	return ans
}

func butterfly(a []int, prm int) {
	g := PrimitiveRoot(prm)
	n := len(a)
	h := CeilPow2(n)

	se := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bits.TrailingZeros(uint(prm - 1))
	e := PowMod(g, (prm-1)>>uint(cnt2), prm)
	ie := InvGcd(e, prm)[1]
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % prm
		ie = ie * ie % prm
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		se[i] = es[i] * now % prm
		now = now * ies[i] % prm
	}
	for ph := 1; ph <= h; ph++ {
		w := 1 << uint(ph-1)
		p := 1 << uint(h-ph)
		now := 1
		for s := 0; s < w; s++ {
			offset := s << uint(h-ph+1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := a[i+offset+p] * now % prm
				a[i+offset] = (l + r) % prm
				a[i+offset+p] = (l - r) % prm
				if a[i+offset+p] < 0 {
					a[i+offset+p] += prm
				}
			}
			now = now * se[bits.TrailingZeros(^uint(s))] % prm
		}
	}
}

func butterflyInv(a []int, prm int) {
	g := PrimitiveRoot(prm)
	n := len(a)
	h := CeilPow2(n)

	sie := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bits.TrailingZeros(uint(prm - 1))
	e := PowMod(g, (prm-1)>>uint(cnt2), prm)
	ie := InvGcd(e, prm)[1]
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % prm
		ie = ie * ie % prm
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		sie[i] = ies[i] * now % prm
		now = now * es[i] % prm
	}
	for ph := h; ph >= 1; ph-- {
		w := 1 << uint(ph-1)
		p := 1 << uint(h-ph)
		inow := 1
		for s := 0; s < w; s++ {
			offset := s << uint(h-ph+1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := a[i+offset+p]
				a[i+offset] = (l + r) % prm
				a[i+offset+p] = (prm + l - r) * inow % prm
			}
			inow = inow * sie[bits.TrailingZeros(^uint(s))] % prm
		}
	}
}
