package mylib

import "sort"

func LowerBound[T Ordered](a []T, x T) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func Lis[T Ordered](a []T) int {
	dp := make([]T, 0, len(a))
	for _, x := range a {
		pos := LowerBound(dp, x)
		if pos == len(dp) {
			dp = append(dp, x)
		} else {
			dp[pos] = x
		}
	}
	return len(dp)
}
