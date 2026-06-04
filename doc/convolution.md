# convolution

Convolution using NTT (Number Theoretic Transform).

## Functions

```go
func Convolution(a, b []int, mod int) []int
func ConvolutionLL(a, b []int) []int
```

`Convolution` computes the convolution of `a` and `b` modulo `mod` using NTT. The modulus must be a prime suitable for NTT (e.g., 998244353, 167772161).

`ConvolutionLL` computes convolution of 64-bit integers using 3-modulus NTT and Garner reconstruction.

## Example

```go
a := []int{1, 2, 3}
b := []int{4, 5, 6}
c := mylib.Convolution(a, b, 998244353)
// c = [4, 13, 28, 27, 18]
```
