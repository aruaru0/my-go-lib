package mylib

func KthStepLoop(N int, K int64, next []int) int {
	pos := 0
	cnt := int64(0)
	seen := make(map[int]int64)
	seen[pos] = 0

	for i := int64(0); i < K; i++ {
		pos = next[pos]
		cnt++
		if _, ok := seen[pos]; ok {
			K -= seen[pos]
			loop := cnt - seen[pos]
			K %= loop
			for j := int64(0); j < K; j++ {
				pos = next[pos]
			}
			return pos
		}
		seen[pos] = cnt
	}
	return pos
}
