package mylib

func Combinations[T any](list []T, choose int) <-chan []T {
	c := make(chan []T)
	go func() {
		defer close(c)
		switch {
		case choose == 0:
			c <- []T{}
		case choose == len(list):
			cp := make([]T, len(list))
			copy(cp, list)
			c <- cp
		case len(list) < choose:
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range Combinations(list[i+1:], choose-1) {
					c <- append([]T{list[i]}, subComb...)
				}
			}
		}
	}()
	return c
}

func generateComb(index []int, s, r int, ch chan []int) {
	if r != 0 {
		if s < 0 {
			return
		}
		generateComb(index, s-1, r, ch)
		index[r-1] = s
		generateComb(index, s-1, r-1, ch)
	} else {
		out := make([]int, len(index))
		copy(out, index)
		ch <- out
	}
}

func CombinationsRec(n, k int) <-chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		if k > n || k < 0 {
			return
		}
		index := make([]int, k)
		generateComb(index, n-1, k, ch)
	}()
	return ch
}
