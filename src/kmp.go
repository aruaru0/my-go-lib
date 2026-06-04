package mylib

type KMP[T comparable] struct {
	pattern []T
	border  []int
}

func NewKMP[T comparable](pattern []T) *KMP[T] {
	kmp := &KMP[T]{pattern: pattern}
	kmp.border = make([]int, len(pattern)+1)
	kmp.border[0] = -1
	j := -1
	for i := 0; i < len(pattern); i++ {
		for j != -1 && pattern[j] != pattern[i] {
			j = kmp.border[j]
		}
		j++
		kmp.border[i+1] = j
	}
	return kmp
}

func (kmp *KMP[T]) Search(text []T) []int {
	if len(kmp.pattern) == 0 {
		res := make([]int, len(text)+1)
		for i := range res {
			res[i] = i
		}
		return res
	}
	var res []int
	j := 0
	for i := 0; i < len(text); i++ {
		for j != -1 && kmp.pattern[j] != text[i] {
			j = kmp.border[j]
		}
		j++
		if j == len(kmp.pattern) {
			res = append(res, i-j+1)
			j = kmp.border[j]
		}
	}
	return res
}
