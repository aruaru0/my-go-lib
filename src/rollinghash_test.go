package mylib

import "testing"

func TestRollingHashDefault(t *testing.T) {
	rh := NewRollingHashDefault("hello")
	if rh.Len() != 5 {
		t.Fatalf("Len() = %d, want 5", rh.Len())
	}
	h := rh.Get(0, 5)
	if h == 0 {
		t.Error("hash of non-empty string should not be 0")
	}
}

func TestRollingHashBasic(t *testing.T) {
	rh := NewRollingHash("abc", 37, int(1e9+7))
	h0_3 := rh.Get(0, 3)
	h0_2 := rh.Get(0, 2)
	h1_3 := rh.Get(1, 3)
	if h0_2 == h0_3 {
		t.Error("different substrings should have different hashes")
	}
	if h0_3 == h1_3 {
		t.Error("different substrings should have different hashes")
	}
}

func TestRollingHashConsistency(t *testing.T) {
	rh1 := NewRollingHash("abcabc", 37, int(1e9+7))
	rh2 := NewRollingHash("abc", 37, int(1e9+7))
	h1 := rh1.Get(0, 3)
	h2 := rh2.Get(0, 3)
	if h1 != h2 {
		t.Errorf("same substring should have same hash: %d vs %d", h1, h2)
	}
}

func TestRollingHashRepeat(t *testing.T) {
	rh := NewRollingHash("aaaa", 37, int(1e9+7))
	seg := rh.Get(0, 2)
	seg2 := rh.Get(2, 4)
	if seg != seg2 {
		t.Errorf("identical substrings should match: %d vs %d", seg, seg2)
	}
}

func TestRollingHashGet(t *testing.T) {
	rh := NewRollingHash("abcde", 37, int(1e9+7))
	for i := 0; i < 5; i++ {
		h := rh.Get(i, i+1)
		if h == 0 {
			t.Errorf("single char hash at %d should not be 0", i)
		}
	}
}

func TestRollingHashEmpty(t *testing.T) {
	rh := NewRollingHashDefault("")
	if rh.Len() != 0 {
		t.Fatal("empty string should have Len 0")
	}
}

func TestRollingHashDifferentStrings(t *testing.T) {
	rh1 := NewRollingHashDefault("abc")
	rh2 := NewRollingHashDefault("xyz")
	if rh1.Get(0, 3) == rh2.Get(0, 3) {
		t.Error("different strings should likely have different hashes")
	}
}
