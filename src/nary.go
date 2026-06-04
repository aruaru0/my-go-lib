package mylib

import "strconv"

func NaryNumbers(N, digits int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		var gen func(string)
		gen = func(s string) {
			if len(s) >= digits {
				ch <- s
				return
			}
			for i := 0; i < N; i++ {
				gen(s + strconv.Itoa(i))
			}
		}
		gen("")
	}()
	return ch
}
