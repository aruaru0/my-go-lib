# fft

Fast Fourier Transform using complex128.

## Functions

```go
func FFT(x []complex128, n int) []complex128
func IFFT(x []complex128, n int) []complex128
```

`FFT` computes the discrete Fourier transform and normalizes by `1/n`. `IFFT` computes the inverse transform.

## Example

```go
n := 8
x := make([]complex128, n)
for i := 0; i < n; i++ {
    x[i] = complex(float64(i+1), 0)
}
f := mylib.FFT(x, n)
inv := mylib.IFFT(f, n)
// inv ≈ x (within floating-point error)
```
