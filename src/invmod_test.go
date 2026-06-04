package mylib

import (
	"testing"
)

func TestInvModTable(t *testing.T) {
	m := 13
	inv := InvModTable(m, m)
	for i := 1; i < m; i++ {
		if got := (i * inv[i]) % m; got != 1 {
			t.Errorf("InvModTable(%d,%d): %d * inv[%d] %% %d = %d, want 1", m, m, i, i, m, got)
		}
	}
}

func TestInvModTableSize(t *testing.T) {
	inv := InvModTable(10, 1000000007)
	if len(inv) != 10 {
		t.Errorf("InvModTable(10, 1e9+7) length = %d, want 10", len(inv))
	}
	if inv[1] != 1 {
		t.Errorf("inv[1] = %d, want 1", inv[1])
	}
}
