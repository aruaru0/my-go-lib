# ncrautogen

nCr 計算のための自動拡張階乗テーブルです。

## 型

```go
type NCRGenerator struct { ... }
```

## 関数

```go
func NewNCRGenerator(mod int) *NCRGenerator
```

## メソッド

```go
func (g *NCRGenerator) NCR(n, r int) int
```

`nCr` を `mod` で割った剰余を計算します。階乗テーブルは必要に応じて自動的に拡張されます。

## 使用例

```go
g := mylib.NewNCRGenerator(1000000007)
c := g.NCR(5, 2) // 10
c = g.NCR(10, 3) // 120
c = g.NCR(100, 50) // 自動的にテーブルが拡張されます
```
