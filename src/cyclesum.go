package mylib

func CycleSum(start, N, mod int, next func(int) int) int {
	visited := make([]int, mod)
	pat := make([]int, 0)
	sum := 0
	cur := start
	for i := 0; i < N; i++ {
		if visited[cur] == 0 {
			visited[cur] = 1
			pat = append(pat, cur)
			sum += cur
		} else {
			startIdx := 0
			for j, v := range pat {
				if v == cur {
					startIdx = j
					break
				}
			}
			loopSum := 0
			prefixSum := 0
			for j := startIdx; j < len(pat); j++ {
				loopSum += pat[j]
			}
			for j := 0; j < startIdx; j++ {
				prefixSum += pat[j]
			}
			remaining := N - i
			loopLen := len(pat) - startIdx
			fullLoops := remaining / loopLen
			sum += prefixSum*fullLoops + loopSum*fullLoops
			rest := remaining % loopLen
			for j := startIdx; j < startIdx+rest; j++ {
				sum += pat[j]
			}
			return sum
		}
		cur = next(cur)
	}
	return sum
}
