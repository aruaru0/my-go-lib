# fft

complex128 を使用した高速フーリエ変換です。

## 関数

```go
func FFT(x []complex128, n int) []complex128
func IFFT(x []complex128, n int) []complex128
```

`FFT` は離散フーリエ変換を計算し、`1/n` で正規化します。`IFFT` は逆変換を計算します。

## 使用例

```go
n := 8
x := make([]complex128, n)
for i := 0; i < n; i++ {
    x[i] = complex(float64(i+1), 0)
}
f := mylib.FFT(x, n)
inv := mylib.IFFT(f, n)
// inv ≈ x (浮動小数点誤差の範囲内)
```
