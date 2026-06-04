# gacha

ガチャ (クーポンコレクター) のコンプリート期待値を求めます。

## 関数

```go
func GachaExpectation(n int) float64
```

## GachaExpectation

クーポンコレクター問題において、`n` 種類すべてのアイテムを集めるために必要な期待試行回数を返します。`n * H_n` として計算されます。ここで `H_n` は n 番目の調和数です。
