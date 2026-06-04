# digitdp

特定の桁の性質を持つ数値をカウントするための Digit DP ユーティリティです。

## 関数

```go
func CountContaining(n int, target int) int64
func CountNonZero(s string, k int) int64
```

## CountContaining

1 から n (両端を含む) までの数値のうち、桁 `target` を含むものの個数をカウントします。

## CountNonZero

1 から s (両端を含む、s は10進数の文字列) までの数値のうち、ちょうど `k` 個のゼロ以外の桁を持つものの個数をカウントします。ABC 154 E に基づきます。
