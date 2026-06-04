package mylib

import (
	"reflect"
	"testing"
)

func TestCompressInt(t *testing.T) {
	tests := []struct {
		s    []int
		want []int
	}{
		{[]int{5, 3, 1, 3, 5}, []int{2, 1, 0, 1, 2}},
		{[]int{10, 20, 10}, []int{0, 1, 0}},
		{[]int{1, 2, 3}, []int{0, 1, 2}},
		{[]int{3, 2, 1}, []int{2, 1, 0}},
		{[]int{1}, []int{0}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		got := Compress(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Compress(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

func TestCompressString(t *testing.T) {
	tests := []struct {
		s    []string
		want []int
	}{
		{[]string{"b", "a", "c"}, []int{1, 0, 2}},
		{[]string{"x", "x", "y"}, []int{0, 0, 1}},
	}
	for _, tt := range tests {
		got := Compress(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Compress(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
