package mylib

import "math/bits"

func CeilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}

func Bsf(n uint) int {
	return bits.TrailingZeros(n)
}
