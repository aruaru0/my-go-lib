package mylib

import (
	"reflect"
	"testing"
)

func TestPrimeFactorMap(t *testing.T) {
	tests := []struct {
		n    int
		want map[int]int
	}{
		{12, map[int]int{2: 2, 3: 1}},
		{7, map[int]int{7: 1}},
		{1, map[int]int{}},
		{100, map[int]int{2: 2, 5: 2}},
	}
	for _, tt := range tests {
		got := PrimeFactorMap(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PrimeFactorMap(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{12, []int{2, 2, 3}},
		{7, []int{7}},
		{1, []int{}},
		{100, []int{2, 2, 5, 5}},
	}
	for _, tt := range tests {
		got := PrimeFactors(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("PrimeFactors(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
