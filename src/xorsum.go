package mylib

func XorSum(dist []int) int64 {
	var ans int64
	N := len(dist)
	for bit := 0; bit < 60; bit++ {
		ones := 0
		for _, v := range dist {
			ones += (v >> bit) & 1
		}
		zeros := N - ones
		ans += int64(ones) * int64(zeros) * int64(1<<bit)
	}
	return ans
}
