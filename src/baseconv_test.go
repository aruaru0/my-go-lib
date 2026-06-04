package mylib

import (
	"reflect"
	"testing"
)

func TestToBaseN(t *testing.T) {
	tests := []struct {
		x, n int
		want []int
	}{
		{0, 2, []int{0}},
		{1, 2, []int{1}},
		{2, 2, []int{0, 1}},
		{5, 2, []int{1, 0, 1}},
		{10, 10, []int{0, 1}},
		{255, 16, []int{15, 15}},
		// negative base -2
		{1, -2, []int{1}},
		{2, -2, []int{0, 1, 1}},
		{3, -2, []int{1, 1, 1}},
		{4, -2, []int{0, 0, 1}},
		{5, -2, []int{1, 0, 1}},
	}
	for _, tt := range tests {
		got := ToBaseN(tt.x, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ToBaseN(%d, %d) = %v, want %v", tt.x, tt.n, got, tt.want)
		}
	}
}

func TestToBaseNInt64(t *testing.T) {
	var x int64 = 10
	var n int64 = 2
	got := ToBaseN(x, n)
	want := []int64{0, 1, 0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ToBaseN[int64](10, 2) = %v, want %v", got, want)
	}
}
