# UnionFind (Disjoint Set Union)

要素を互いに重複しない複数のグループに分割して管理し、「2つの要素が同じグループに属するか判定する（Same）」「2つのグループを併合する（Merge）」というクエリを高速に行うデータ構造です（Union-Find）。

## 主なユースケース（どういう時に使うか）

- 無向グラフの連結成分数カウント
- Kruskal法による最小全域木の構築
- グループ所属の高速判定

## なぜ使うのか（メリット）

- 経路圧縮とサイズによる併合ルールを用いることで、各クエリをほぼ $O(1)$ （逆アッカーマン関数 $lpha(N)$ 時間）で実行できる極めて強力なデータ構造です。

---

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