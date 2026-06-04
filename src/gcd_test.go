package mylib

import "testing"

func TestGcd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{12, 8, 4},
		{7, 13, 1},
		{0, 5, 5},
		{6, 0, 6},
		{100, 25, 25},
		{1, 1, 1},
	}
	for _, tt := range tests {
		got := Gcd(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Gcd(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGcdInt64(t *testing.T) {
	var a, b int64 = 48, 18
	got := Gcd(a, b)
	if got != 6 {
		t.Errorf("Gcd[int64](%d, %d) = %d, want %d", a, b, got, 6)
	}
}

func TestGcdUint(t *testing.T) {
	var a, b uint = 36, 24
	got := Gcd(a, b)
	if got != 12 {
		t.Errorf("Gcd[uint](%d, %d) = %d, want %d", a, b, got, 12)
	}
}

func TestLcm(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{12, 8, 24},
		{7, 13, 91},
		{4, 6, 12},
		{5, 5, 5},
		{1, 10, 10},
	}
	for _, tt := range tests {
		got := Lcm(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Lcm(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLcmInt64(t *testing.T) {
	var a, b int64 = 6, 8
	got := Lcm(a, b)
	if got != 24 {
		t.Errorf("Lcm[int64](%d, %d) = %d, want %d", a, b, got, 24)
	}
}
