package mylib

func Subsets(i int) []int {
	var res []int
	for j := i; j != 0; j = (j - 1) & i {
		res = append(res, j)
	}
	return res
}
