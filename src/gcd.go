package mylib

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Gcd[T Integer](a, b T) T {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Lcm[T Integer](a, b T) T {
	return a / Gcd(a, b) * b
}
