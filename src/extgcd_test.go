package mylib

import (
	"testing"
)

func TestExtGcd(t *testing.T) {
	g, x, y := ExtGcd(6, 15)
	if g != 3 {
		t.Errorf("ExtGcd(6,15): g = %d, want 3", g)
	}
	if 6*x+15*y != g {
		t.Errorf("ExtGcd(6,15): 6*%d + 15*%d = %d, want %d", x, y, 6*x+15*y, g)
	}
}

func TestExtGcdCoprime(t *testing.T) {
	g, x, y := ExtGcd(7, 11)
	if g != 1 {
		t.Errorf("ExtGcd(7,11): g = %d, want 1", g)
	}
	if 7*x+11*y != 1 {
		t.Errorf("ExtGcd(7,11): 7*%d + 11*%d = %d, want 1", x, y, 7*x+11*y)
	}
}

func TestExtGcdGeneric(t *testing.T) {
	g, x, y := ExtGcd(int64(24), int64(36))
	if g != 12 {
		t.Errorf("ExtGcd(24,36): g = %d, want 12", g)
	}
	if 24*x+36*y != g {
		t.Errorf("ExtGcd(24,36): 24*%d + 36*%d = %d, want %d", x, y, 24*x+36*y, g)
	}
}
