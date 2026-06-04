# divsubstr

10 進数文字列の中で P で割り切れる部分文字列の個数を数えます。

## 関数

```go
func CountDivisibleSubstrings(s string, p int) int
```

## CountDivisibleSubstrings

`s` の部分文字列のうち、`p` で割り切れるものの個数をカウントします。右から左への mod 累積と、マップを使用した剰余の追跡を行います。`p = 2` または `p = 5` の場合は、特別に最適化された経路を使用します。
