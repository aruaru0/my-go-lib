package mylib

func SlideMin[T Ordered](a []T, k int) []T {
	n := len(a)
	b := make([]T, n-k+1)
	s, t := 0, 0
	deq := make([]int, n)
	for i := 0; i < n; i++ {
		for s < t && a[deq[t-1]] >= a[i] {
			t--
		}
		deq[t] = i
		t++
		if i-k+1 >= 0 {
			b[i-k+1] = a[deq[s]]
			if deq[s] == i-k+1 {
				s++
			}
		}
	}
	return b
}

func SlideMax[T Ordered](a []T, k int) []T {
	n := len(a)
	b := make([]T, n-k+1)
	s, t := 0, 0
	deq := make([]int, n)
	for i := 0; i < n; i++ {
		for s < t && a[deq[t-1]] <= a[i] {
			t--
		}
		deq[t] = i
		t++
		if i-k+1 >= 0 {
			b[i-k+1] = a[deq[s]]
			if deq[s] == i-k+1 {
				s++
			}
		}
	}
	return b
}
