package mylib

func GachaExpectation(n int) float64 {
	var h float64
	for i := 1; i <= n; i++ {
		h += 1.0 / float64(i)
	}
	return float64(n) * h
}
