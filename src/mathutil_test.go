package mylib

import (
	"testing"
)

func TestFloorSum(t *testing.T) {
	if got := FloorSum(4, 3, 2, 1); got != 4 {
		t.Errorf("FloorSum(4,3,2,1) = %d, want 4", got)
	}
}

func TestPowMod(t *testing.T) {
	if got := PowMod(2, 10, 1000); got != 24 {
		t.Errorf("PowMod(2,10,1000) = %d, want 24", got)
	}
	if got := PowMod(3, 0, 100); got != 1 {
		t.Errorf("PowMod(3,0,100) = %d, want 1", got)
	}
	if got := PowMod(5, 100, 1); got != 0 {
		t.Errorf("PowMod(5,100,1) = %d, want 0", got)
	}
}

func TestSafeMod(t *testing.T) {
	if got := SafeMod(-5, 3); got != 1 {
		t.Errorf("SafeMod(-5,3) = %d, want 1", got)
	}
	if got := SafeMod(7, 3); got != 1 {
		t.Errorf("SafeMod(7,3) = %d, want 1", got)
	}
}

func TestInvGcd(t *testing.T) {
	res := InvGcd(3, 10)
	if res[0] != 1 || res[1] != 7 {
		t.Errorf("InvGcd(3,10) = %v, want [1 7]", res)
	}
}

func TestInvMod(t *testing.T) {
	inv := InvMod(3, 10)
	if got := (3 * inv) % 10; got != 1 {
		t.Errorf("InvMod(3,10) = %d, 3*inv%%10 = %d, want 1", inv, got)
	}
}

func TestCrt(t *testing.T) {
	res := Crt([]int{2, 3}, []int{3, 5})
	if res[0] != 8 || res[1] != 15 {
		t.Errorf("Crt([2,3],[3,5]) = %v, want [8 15]", res)
	}
}

func TestPrimitiveRoot(t *testing.T) {
	if got := PrimitiveRoot(998244353); got != 3 {
		t.Errorf("PrimitiveRoot(998244353) = %d, want 3", got)
	}
	if got := PrimitiveRoot(2); got != 1 {
		t.Errorf("PrimitiveRoot(2) = %d, want 1", got)
	}
}
