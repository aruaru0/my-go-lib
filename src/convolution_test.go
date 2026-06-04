package mylib

import (
	"testing"
)

func TestConvolutionSmall(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	mod := 998244353
	got := Convolution(a, b, mod)
	want := []int{4, 13, 28, 27, 18}
	if len(got) != len(want) {
		t.Fatalf("Convolution length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Convolution[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestConvolutionEmpty(t *testing.T) {
	got := Convolution([]int{}, []int{1, 2}, 998244353)
	if len(got) != 0 {
		t.Errorf("Convolution empty = %v, want empty", got)
	}
}

func TestConvolutionNaiveBranch(t *testing.T) {
	a := []int{1}
	b := []int{2}
	mod := 998244353
	got := Convolution(a, b, mod)
	if got[0] != 2 {
		t.Errorf("Convolution([1],[2]) = %d, want 2", got[0])
	}
}

func TestConvolutionLL(t *testing.T) {
	a := []int{1000000000, 2000000000}
	b := []int{3000000000, 4000000000}
	got := ConvolutionLL(a, b)
	want := []int{3000000000, 10000000000, 8000000000000000000}
	if len(got) != len(want) {
		t.Fatalf("ConvolutionLL length = %d, want %d", len(got), len(want))
	}
}
