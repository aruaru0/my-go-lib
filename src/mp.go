package mylib

type MP struct {
	pattern string
	border  []int
}

func NewMP(pattern string) *MP {
	mp := &MP{pattern: pattern}
	mp.border = make([]int, len(pattern)+1)
	mp.border[0] = -1
	j := -1
	for i := 0; i < len(pattern); i++ {
		for j != -1 && pattern[j] != pattern[i] {
			j = mp.border[j]
		}
		j++
		mp.border[i+1] = j
	}
	return mp
}

func (mp *MP) FindAll(text string) []int {
	if len(mp.pattern) == 0 {
		res := make([]int, len(text)+1)
		for i := range res {
			res[i] = i
		}
		return res
	}
	var res []int
	j := 0
	for i := 0; i < len(text); i++ {
		for j != -1 && mp.pattern[j] != text[i] {
			j = mp.border[j]
		}
		j++
		if j == len(mp.pattern) {
			res = append(res, i-j+1)
			j = mp.border[j]
		}
	}
	return res
}
