package mylib

func ToBaseN[T Integer](x T, n T) []T {
	var ret []T
	for x != 0 {
		r := x % n
		if r < 0 {
			r += -n
		}
		x = (x - r) / n
		ret = append(ret, r)
	}
	if len(ret) == 0 {
		ret = append(ret, 0)
	}
	return ret
}
