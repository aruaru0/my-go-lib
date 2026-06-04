package mylib

func MaxGainK(i, N, K int, p, c []int) int {
	MAX := int(-1e18)
	used := make([]bool, N)
	pos := i
	cost := 0
	rt := []int{i}
	tot := []int{0}
	for !used[pos] {
		used[pos] = true
		pos = p[pos] - 1
		cost += c[pos]
		rt = append(rt, pos)
		tot = append(tot, cost)
	}
	if K < len(tot) {
		for j := 1; j <= K; j++ {
			MAX = max(MAX, tot[j])
		}
	} else if cost < 0 {
		for _, e := range tot {
			MAX = max(MAX, e)
		}
	} else {
		loop := K / (len(tot) - 1)
		if loop > 0 {
			loop--
			sum := loop * tot[len(tot)-1]
			for _, e := range tot {
				MAX = max(MAX, sum+e)
			}
		}
		loop = K / (len(tot) - 1)
		sum := loop * tot[len(tot)-1]
		rest := K % (len(tot) - 1)
		MAX = max(MAX, sum)
		for i := 1; i <= rest; i++ {
			MAX = max(MAX, sum+tot[i])
		}
	}
	return MAX
}
