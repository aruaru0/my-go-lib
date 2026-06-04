package mylib

import (
	"reflect"
	"testing"
)

func TestSuffixArrayString(t *testing.T) {
	sa := SuffixArrayString("banana")
	// Expected: [5 3 1 0 4 2] for "banana"
	// Verify by checking that s[sa[i]:] is sorted
	if len(sa) != 6 {
		t.Fatalf("SuffixArrayString(banana) length = %d, want 6", len(sa))
	}
	s := "banana"
	for i := 0; i < len(sa)-1; i++ {
		if s[sa[i]:] >= s[sa[i+1]:] {
			t.Errorf("suffix at %d (%q) should be < suffix at %d (%q)", sa[i], s[sa[i]:], sa[i+1], s[sa[i+1]:])
		}
	}
}

func TestSuffixArrayEmpty(t *testing.T) {
	sa := SuffixArray(nil, 0)
	if len(sa) != 0 {
		t.Errorf("expected empty, got %v", sa)
	}
}

func TestSuffixArraySingle(t *testing.T) {
	sa := SuffixArray([]int{0}, 1)
	want := []int{0}
	if !reflect.DeepEqual(sa, want) {
		t.Errorf("SuffixArray([0]) = %v, want %v", sa, want)
	}
}

func TestSuffixArraySortedProperty(t *testing.T) {
	s := []int{1, 2, 2, 1, 3}
	sa := SuffixArrayInt(s)
	if len(sa) != len(s) {
		t.Fatalf("SuffixArrayInt(%v) length = %d, want %d", s, len(sa), len(s))
	}
	for i := 0; i < len(sa)-1; i++ {
		a, b := s[sa[i]:], s[sa[i+1]:]
		if !sliceLess(a, b) {
			t.Errorf("suffix[%d] %v should be < suffix[%d] %v", sa[i], a, sa[i+1], b)
		}
	}
}

func TestSuffixArrayUnique(t *testing.T) {
	s := []int{5, 1, 3, 2, 4}
	sa := SuffixArrayInt(s)
	if len(sa) != len(s) {
		t.Fatalf("SuffixArrayInt(%v) length = %d, want %d", s, len(sa), len(s))
	}
	for i := 0; i < len(sa)-1; i++ {
		a, b := s[sa[i]:], s[sa[i+1]:]
		if !sliceLess(a, b) {
			t.Errorf("suffix[%d] %v should be < suffix[%d] %v", sa[i], a, sa[i+1], b)
		}
	}
}

func sliceLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func TestLcpArray(t *testing.T) {
	sa := SuffixArrayString("banana")
	lcp := LcpArrayString("banana", sa)
	// LCP array for "banana": [1 3 0 0 2]
	if len(lcp) != 5 {
		t.Fatalf("LcpArray length = %d, want 5", len(lcp))
	}
}
