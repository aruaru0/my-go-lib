# bits

ビット操作ユーティリティです。

## 関数

```go
func CeilPow2(n int) int
func Bsf(n uint) int
```

## CeilPow2

`2^x >= n` を満たす最小の `x` を返します。

## Bsf

前方ビットスキャンです。`n` の末尾のゼロビットの数を返します (`bits.TrailingZeros` を使用します)。
