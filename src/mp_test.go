package mylib

import (
	"reflect"
	"testing"
)

func TestMPFindAll(t *testing.T) {
	tests := []struct {
		pattern string
		text    string
		want    []int
	}{
		{"abc", "abcabcabc", []int{0, 3, 6}},
		{"abc", "abababc", []int{4}},
		{"aaa", "aaaaa", []int{0, 1, 2}},
		{"a", "abcabc", []int{0, 3}},
		{"abc", "def", nil},
		{"", "", []int{0}},
		{"", "abc", []int{0, 1, 2, 3}},
		{"abc", "", nil},
		{"x", "x", []int{0}},
		{"ab", "ababab", []int{0, 2, 4}},
	}
	for _, tt := range tests {
		mp := NewMP(tt.pattern)
		got := mp.FindAll(tt.text)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewMP(%q).FindAll(%q) = %v, want %v", tt.pattern, tt.text, got, tt.want)
		}
	}
}

func TestMPNoMatch(t *testing.T) {
	mp := NewMP("xyz")
	got := mp.FindAll("abc")
	if len(got) != 0 {
		t.Errorf("expected no matches, got %v", got)
	}
}
