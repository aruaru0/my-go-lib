# KthStepLoop — ループ検出付き K ステップグラフ探索

## 関数

```go
func KthStepLoop(N int, K int64, next []int) int
```

KthStepLoop は、ノード 0 から開始して正確に K ステップ後に到達するノードを返します。KthStep とは異なり、マップベースのループ検出を使用するため、K が非常に大きくグラフにサイクルが含まれる場合に効率的です。
