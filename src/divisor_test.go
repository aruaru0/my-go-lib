package mylib

import (
	"sort"
	"testing"
)

func TestDivisors(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{12, []int{1, 2, 3, 4, 6, 12}},
		{7, []int{1, 7}},
		{1, []int{1}},
		{16, []int{1, 2, 4, 8, 16}},
		{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
	}
	for _, tt := range tests {
		got := Divisors(tt.n)
		sort.Ints(got)
		if !sameElements(got, tt.want) {
			t.Errorf("Divisors(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func sameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
