package mylib

import "sort"

func Compress[T Ordered](s []T) []int {
	uniq := make([]T, len(s))
	copy(uniq, s)
	sort.Slice(uniq, func(i, j int) bool { return uniq[i] < uniq[j] })
	j := 0
	for i := 0; i < len(uniq); i++ {
		if i == 0 || uniq[i] != uniq[i-1] {
			uniq[j] = uniq[i]
			j++
		}
	}
	uniq = uniq[:j]
	m := make(map[T]int)
	for i, v := range uniq {
		m[v] = i
	}
	res := make([]int, len(s))
	for i, v := range s {
		res[i] = m[v]
	}
	return res
}
