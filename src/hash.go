package mylib

func HashInts[T Integer](x []T) int {
	ret := 0
	for _, e := range x {
		ret = ret*37 + int(e)
		ret %= 1e9 + 7
	}
	return ret
}
