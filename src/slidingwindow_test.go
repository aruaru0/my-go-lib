package mylib

import (
	"reflect"
	"testing"
)

func TestSlideMin(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{-1, -3, -3, -3, 3, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{1, 2, 3, 4}},
		{[]int{6, 5, 4, 3, 2, 1}, 3, []int{4, 3, 2, 1}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1}},
	}
	for _, tt := range tests {
		got := SlideMin(tt.a, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SlideMin(%v, %d) = %v, want %v", tt.a, tt.k, got, tt.want)
		}
	}
}

func TestSlideMax(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, 3, []int{6, 5, 4, 3}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for _, tt := range tests {
		got := SlideMax(tt.a, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SlideMax(%v, %d) = %v, want %v", tt.a, tt.k, got, tt.want)
		}
	}
}
