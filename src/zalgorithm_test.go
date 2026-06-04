package mylib

import (
	"reflect"
	"testing"
)

func TestZAlgorithmString(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{"aaaaa", []int{5, 4, 3, 2, 1}},
		{"abcabcabc", []int{9, 0, 0, 6, 0, 0, 3, 0, 0}},
		{"abcde", []int{5, 0, 0, 0, 0}},
		{"", []int{}},
		{"a", []int{1}},
		{"aa", []int{2, 1}},
		{"abacabab", []int{8, 0, 1, 0, 3, 0, 2, 0}},
	}
	for _, tt := range tests {
		got := ZAlgorithmString(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ZAlgorithmString(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

func TestZAlgorithmInt(t *testing.T) {
	s := []int{1, 2, 1, 2, 1}
	z := ZAlgorithm(s)
	// expected: length 5
	if len(z) != 5 {
		t.Fatalf("ZAlgorithm(%v) length = %d, want 5", s, len(z))
	}
	if z[0] != 5 {
		t.Errorf("z[0] = %d, want 5", z[0])
	}
}

func TestZAlgorithmEmpty(t *testing.T) {
	got := ZAlgorithm([]int{})
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}
