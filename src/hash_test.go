package mylib

import "testing"

func hashWant(seq ...int) int {
	r := 0
	for _, e := range seq {
		r = r*37 + e
		r %= 1000000007
	}
	return r
}

func TestHashInts(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1}, hashWant(1)},
		{[]int{1, 2}, hashWant(1, 2)},
		{[]int{1, 2, 3}, hashWant(1, 2, 3)},
	}
	for _, tt := range tests {
		got := HashInts(tt.in)
		if got != tt.want {
			t.Errorf("HashInts(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestHashIntsConsistency(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	if HashInts(a) != HashInts(b) {
		t.Error("HashInts not consistent")
	}
}

func TestHashIntsInt64(t *testing.T) {
	in := []int64{10, 20, 30}
	got := HashInts(in)
	want := hashWant(10, 20, 30)
	if got != want {
		t.Errorf("HashInts(%v) = %d, want %d", in, got, want)
	}
}

func TestHashIntsUint(t *testing.T) {
	in := []uint{5, 15, 25}
	got := HashInts(in)
	if got == 0 && len(in) > 0 {
		t.Error("HashInts should not return 0 for non-empty uint slice")
	}
}
