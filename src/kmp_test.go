package mylib

import (
	"reflect"
	"testing"
)

func TestKMPInt(t *testing.T) {
	tests := []struct {
		pattern []int
		text    []int
		want    []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3, 1, 2, 3}, []int{0, 3}},
		{[]int{1, 2}, []int{1, 2, 1, 2, 1, 2}, []int{0, 2, 4}},
		{[]int{4, 5}, []int{1, 2, 3}, nil},
		{[]int{1}, []int{1, 1, 1}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		kmp := NewKMP(tt.pattern)
		got := kmp.Search(tt.text)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewKMP(%v).Search(%v) = %v, want %v", tt.pattern, tt.text, got, tt.want)
		}
	}
}

func TestKMPString(t *testing.T) {
	kmp := NewKMP([]byte("abc"))
	got := kmp.Search([]byte("abcabcabc"))
	want := []int{0, 3, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewKMP(abc).Search(abcabcabc) = %v, want %v", got, want)
	}
}

func TestKMPEmptyPattern(t *testing.T) {
	kmp := NewKMP([]int{})
	got := kmp.Search([]int{1, 2, 3})
	want := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("empty pattern should match at every position, got %v, want %v", got, want)
	}
}
