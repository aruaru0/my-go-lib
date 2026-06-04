# UnionFind (Disjoint Set Union)

`UnionFind` は、集合を互いに素な部分集合に分割して管理するデータ構造です。次の操作をサポートします:

- **Unite**: 2 つの部分集合を 1 つに併合
- **Same**: 2 つの要素が同じ部分集合に属するか確認
- **Size**: 要素が属する部分集合のサイズを取得

## コンストラクタ

```go
func NewUnionFind(N int) *UnionFind
```

`N` 個の要素 (0..N-1) を持つ新しい UnionFind を作成します。各要素は最初は自身のみからなる集合に属します。

## メソッド

```go
func (p *UnionFind) Unite(x, y int)
```
`x` と `y` を含む集合を併合します。既に同じ集合に属している場合は何もしません。

```go
func (p *UnionFind) Same(x, y int) bool
```
`x` と `y` が同じ集合に属する場合は `true` を返します。

```go
func (p *UnionFind) Size(x int) int
```
`x` を含む集合の要素数を返します。

## 使用例

```go
uf := NewUnionFind(10)
uf.Unite(0, 1)
uf.Unite(1, 2)
fmt.Println(uf.Same(0, 2)) // true
fmt.Println(uf.Size(0))    // 3
```

## 計算量

`Unite` と `Same` はどちらも、経路圧縮とサイズによる併合により、ならし O(1) 時間 (逆アッカーマン関数) で動作します。
