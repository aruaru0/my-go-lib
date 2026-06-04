# convolution

NTT (Number Theoretic Transform) を使用した畳み込みです。

## 関数

```go
func Convolution(a, b []int, mod int) []int
func ConvolutionLL(a, b []int) []int
```

`Convolution` は、NTT を使用して `a` と `b` の畳み込みを `mod` で割った剰余で計算します。法は NTT に適した素数である必要があります (例: 998244353, 167772161)。

`ConvolutionLL` は、3 元 NTT と Garner の再構成を使用して 64 ビット整数の畳み込みを計算します。

## 使用例

```go
a := []int{1, 2, 3}
b := []int{4, 5, 6}
c := mylib.Convolution(a, b, 998244353)
// c = [4, 13, 28, 27, 18]
```
