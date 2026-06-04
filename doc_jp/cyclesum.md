# cyclesum

mod に基づくサイクル検出を用いた、周期系列の効率的な総和計算を行います。

## 関数

```go
func CycleSum(start, N, mod int, next func(int) int) int
```

## CycleSum

`next` で定義される数列の最初の `N` 項の総和を、`start` から開始して計算します。繰り返される mod の剰余に基づいてサイクルを検出し、そのサイクルを利用して N の値にかかわらず O(mod) 時間で総和を計算します。
