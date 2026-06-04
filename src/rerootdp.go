package mylib

type distInfo [2]int

func TreeDistances(N int, node [][]int) []int {
	dp1 := make([]int, N)
	dp2 := make([]int, N)
	ans := make([]int, N)

	var dfs1 func(v, parent int)
	dfs1 = func(v, parent int) {
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			dfs1(to, v)
			if dp1[v] < dp1[to]+1 {
				dp1[v] = dp1[to] + 1
			}
		}
	}

	var dfs2 func(v, parent int)
	dfs2 = func(v, parent int) {
		firstMax, secondMax := -1, -1
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			length := dp1[to] + 1
			if length > firstMax {
				secondMax = firstMax
				firstMax = length
			} else if length > secondMax {
				secondMax = length
			}
		}
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			use := firstMax
			if dp1[to]+1 == firstMax {
				use = secondMax
			}
			dp2[to] = max(dp2[v]+1, use+1)
			dfs2(to, v)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	for i := 0; i < N; i++ {
		ans[i] = max(dp1[i], dp2[i])
	}
	return ans
}

func TreeDistancesWithNode(N int, node [][]int) ([]int, []int) {
	dp1 := make([]distInfo, N)
	dp2 := make([]distInfo, N)

	ansDist := make([]int, N)
	ansIdx := make([]int, N)

	var dfs1 func(v, parent int)
	dfs1 = func(v, parent int) {
		dp1[v] = distInfo{0, v}
		for _, to := range node[v] {
			if to == parent {
				continue
			}
			dfs1(to, v)
			dist := dp1[to][0] + 1
			idx := dp1[to][1]

			if dp1[v][0] < dist {
				dp1[v][0] = dist
				dp1[v][1] = idx
			} else if dp1[v][0] == dist {
				dp1[v][1] = max(dp1[v][1], idx)
			}
		}
	}

	var dfs2 func(v, parent int)
	dfs2 = func(v, parent int) {
		first := distInfo{-1, -1}
		second := distInfo{-1, -1}

		for _, to := range node[v] {
			if to == parent {
				continue
			}
			info := distInfo{dp1[to][0] + 1, dp1[to][1]}

			if info[0] > first[0] || (info[0] == first[0] && info[1] > first[1]) {
				second = first
				first = info
			} else if info[0] > second[0] || (info[0] == second[0] && info[1] > second[1]) {
				second = info
			}
		}

		for _, to := range node[v] {
			if to == parent {
				continue
			}

			parentInfo := distInfo{dp2[v][0] + 1, dp2[v][1]}

			siblingInfo := first
			if first[0] == dp1[to][0]+1 && first[1] == dp1[to][1] {
				siblingInfo = second
			}
			if siblingInfo[0] != -1 {
				siblingInfo[0]++
			}

			if parentInfo[0] > siblingInfo[0] {
				dp2[to] = parentInfo
			} else if parentInfo[0] < siblingInfo[0] {
				dp2[to] = siblingInfo
			} else {
				dp2[to] = distInfo{parentInfo[0], max(parentInfo[1], siblingInfo[1])}
			}

			dfs2(to, v)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	for i := 0; i < N; i++ {
		info1 := dp1[i]
		info2 := dp2[i]

		if info1[0] > info2[0] {
			ansDist[i] = info1[0]
			ansIdx[i] = info1[1]
		} else if info1[0] < info2[0] {
			ansDist[i] = info2[0]
			ansIdx[i] = info2[1]
		} else {
			ansDist[i] = info1[0]
			ansIdx[i] = max(info1[1], info2[1])
		}
	}
	return ansDist, ansIdx
}
