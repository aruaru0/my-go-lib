package mylib

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		in   []int
		want bool
		out  []int
	}{
		{[]int{1, 2, 3}, true, []int{1, 3, 2}},
		{[]int{3, 2, 1}, false, []int{3, 2, 1}},
		{[]int{1, 1, 3}, true, []int{1, 3, 1}},
		{[]int{1}, false, []int{1}},
		{[]int{}, false, []int{}},
		{[]int{1, 3, 2}, true, []int{2, 1, 3}},
		{[]int{2, 3, 1}, true, []int{3, 1, 2}},
	}
	for _, tt := range tests {
		x := make([]int, len(tt.in))
		copy(x, tt.in)
		got := NextPermutation(x)
		if got != tt.want || !reflect.DeepEqual(x, tt.out) {
			t.Errorf("NextPermutation(%v) = %v, result = %v, want result = %v", tt.in, got, x, tt.out)
		}
	}
}
