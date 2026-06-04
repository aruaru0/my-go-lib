# median

追加・削除操作に対応した動的中央値データ構造です。

## 型

```go
type Median struct { ... }
func NewMedian() *Median
func (m *Median) Add(x int)
func (m *Median) Remove(x int)
func (m *Median) Median() int
```

## Median

整数のマルチセットを管理し、中央値のクエリをサポートします。2 つのヒープ (下半分は最大ヒープ、上半分は最小ヒープ) で実装されています。要素数が偶数の場合は、中央の 2 つの値の平均の floor を返します。
