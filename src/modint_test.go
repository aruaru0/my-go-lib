package mylib

import (
	"testing"
)

func TestModintAdd(t *testing.T) {
	m := NewModint(1000000007)
	if got := m.Add(10, 20); got != 30 {
		t.Errorf("Add(10,20) = %d, want 30", got)
	}
	if got := m.Add(1000000000, 100); got != 93 {
		t.Errorf("Add(1000000000,100) = %d, want 93", got)
	}
}

func TestModintSub(t *testing.T) {
	m := NewModint(1000000007)
	if got := m.Sub(20, 10); got != 10 {
		t.Errorf("Sub(20,10) = %d, want 10", got)
	}
	if got := m.Sub(10, 20); got != 1000000007-10 {
		t.Errorf("Sub(10,20) = %d, want %d", got, 1000000007-10)
	}
}

func TestModintMul(t *testing.T) {
	m := NewModint(1000000007)
	if got := m.Mul(1000000, 1000000); got != 999993007 {
		t.Errorf("Mul(1000000,1000000) = %d, want 999993007", got)
	}
}

func TestModintPow(t *testing.T) {
	m := NewModint(1000000007)
	if got := m.Pow(2, 10); got != 1024 {
		t.Errorf("Pow(2,10) = %d, want 1024", got)
	}
	if got := m.Pow(2, 0); got != 1 {
		t.Errorf("Pow(2,0) = %d, want 1", got)
	}
	if got := m.Pow(3, 100); got != 886041711 {
		t.Errorf("Pow(3,100) = %d, want 886041711", got)
	}
}

func TestModintInv(t *testing.T) {
	m := NewModint(1000000007)
	inv2 := m.Inv(2)
	if got := m.Mul(2, inv2); got != 1 {
		t.Errorf("2 * inv(2) = %d, want 1", got)
	}
}

func TestModintDiv(t *testing.T) {
	m := NewModint(1000000007)
	got := m.Div(10, 2)
	if got != 5 {
		t.Errorf("Div(10,2) = %d, want 5", got)
	}
}

func TestModintPowMatrix(t *testing.T) {
	m := NewModint(1000000007)
	A := [][]int{{1, 1}, {1, 0}}
	P := m.PowMatrix(A, 10)
	if P[0][0] != 89 || P[0][1] != 55 || P[1][0] != 55 || P[1][1] != 34 {
		t.Errorf("PowMatrix(A,10) = %v, want [[89 55] [55 34]]", P)
	}
}
