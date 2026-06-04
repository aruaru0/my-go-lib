package mylib

func ExtGcd[T Integer](a, b T) (T, T, T) {
	if b == 0 {
		return a, 1, 0
	}
	d, y, x := ExtGcd(b, a%b)
	y -= a / b * x
	return d, x, y
}
