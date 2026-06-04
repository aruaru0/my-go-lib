# KthStep — ダブリングによる K ステップグラフ探索

## 関数

```go
func KthStep(N int, K int64, next []int) int
```

KthStep は、ノード 0 から開始して正確に K ステップ後に到達するノードを返します。`next[i]` はノード i から 1 ステップでの移動先です。

60 ビットテーブルを使用した Binary Lifting (ダブリング) を使用します。
