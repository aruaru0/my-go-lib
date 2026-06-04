package mylib

import (
	"math"
	"math/cmplx"
	"testing"
)

func TestFFTRoundtrip(t *testing.T) {
	n := 8
	x := make([]complex128, n)
	for i := 0; i < n; i++ {
		x[i] = complex(float64(i+1), 0)
	}
	f := FFT(x, n)
	inv := IFFT(f, n)
	for i := 0; i < n; i++ {
		if math.Abs(real(inv[i])-real(x[i])) > 1e-10 || math.Abs(imag(inv[i])-imag(x[i])) > 1e-10 {
			t.Errorf("FFT/IFFT roundtrip failed at %d: got %v, want %v", i, inv[i], x[i])
		}
	}
}

func TestFFTSine(t *testing.T) {
	n := 4
	x := make([]complex128, n)
	x[0] = complex(1, 0)
	x[1] = complex(2, 0)
	x[2] = complex(3, 0)
	x[3] = complex(4, 0)
	f := FFT(x, n)
	inv := IFFT(f, n)
	for i := 0; i < n; i++ {
		if cmplx.Abs(inv[i]-x[i]) > 1e-10 {
			t.Errorf("Roundtrip at %d: %v != %v", i, inv[i], x[i])
		}
	}
}

func TestIFFTDirect(t *testing.T) {
	n := 1
	x := []complex128{complex(5, 0)}
	inv := IFFT(x, n)
	if real(inv[0]) != 5 || imag(inv[0]) != 0 {
		t.Errorf("IFFT of [5] = %v, want 5", inv[0])
	}
}
