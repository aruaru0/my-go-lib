# DSU (Disjoint Set Union)

`DSU` は、集合を互いに素な部分集合に分割して管理するデータ構造です。Union-Find とも呼ばれます。経路圧縮とサイズによる併合をサポートします。

## コンストラクタ

```go
func NewDsu(n int) *DSU
```

`n` 個の要素 (0..n-1) を持つ新しい DSU を作成します。各要素は最初は自身のみからなる集合に属します。

## メソッド

```go
func (d *DSU) Merge(a, b int) int
```
`a` と `b` を含む集合を併合します。新しいリーダーを返します。

```go
func (d *DSU) Same(a, b int) bool
```
`a` と `b` が同じ集合に属する場合は `true` を返します。

```go
func (d *DSU) Leader(a int) int
```
`a` を含む集合の代表要素を返します。

```go
func (d *DSU) Size(a int) int
```
`a` を含む集合の要素数を返します。

```go
func (d *DSU) Groups() [][]int
```
すべてのグループ (連結成分) のリストを返します。

## 使用例

```go
d := NewDsu(6)
d.Merge(0, 1)
d.Merge(2, 3)
d.Merge(0, 2)
fmt.Println(d.Same(1, 3)) // true
fmt.Println(d.Size(0))    // 4
fmt.Println(d.Groups())   // [[0 1 2 3] [4] [5]]
```

## 計算量

`Merge` と `Same` はどちらもならし O(1) 時間 (逆アッカーマン関数) で動作します。
