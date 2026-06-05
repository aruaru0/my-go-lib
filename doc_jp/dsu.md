# DSU (Disjoint Set Union)

要素を互いに重複しない複数のグループに分割して管理し、「2つの要素が同じグループに属するか判定する（Same）」「2つのグループを併合する（Merge）」というクエリを高速に行うデータ構造です（Union-Find）。

## 主なユースケース（どういう時に使うか）

- **グラフの連結成分の管理（無向グラフの連結判定）**
- **最小全域木（Kruskal法）の構築**
- グループ分けの動的更新やクラスタリングの追跡

## なぜ使うのか（メリット）

- 経路圧縮（Path Compression）とサイズによる併合（Union by Size）を併用することで、各操作をほぼ定数時間であるならし $O(\alpha(N))$ （$\alpha$ は逆アッカーマン関数で、実用上は $4$ 以下）で超高速に処理できます。

---

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