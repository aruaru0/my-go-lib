# LCA — 最小共通祖先 (Binary Lifting)

根付き木において、指定された2つの頂点 $u$ と $v$ の共通の祖先の中で、最も根から遠い（最も深い）頂点を求めるアルゴリズムです（ダブリングによる実装）。

## 主なユースケース（どういう時に使うか）

- 木における2頂点間の最短パスや距離の算出
- 組織図やファイルシステムのような階層構造における共通結合点の特定

## なぜ使うのか（メリット）

- $O(N \log N)$ の前処理（ダブリングテーブルの構築）をしておくことで、任意の2頂点に対する共通祖先クエリに $O(\log N)$ で高速に応答できます。

---

## 型

```go
type LCA struct { ... }
```

## 関数

```go
func NewLCA(n int) *LCA
func (l *LCA) AddEdge(from, to int)
func (l *LCA) Build(root int)
func (l *LCA) LCA(u, v int) int
func (l *LCA) Dist(u, v int) int
```

LCA は、Binary Lifting (ダブリング) を使用して根付き木上の最小共通祖先クエリを提供します。

- NewLCA は n 個のノードからなる木を作成します。
- AddEdge は無向辺を追加します。
- Build は祖先テーブルを事前計算します (すべての辺を追加した後、クエリの前に呼び出す必要があります)。
- LCA は u と v の LCA を返します。
- Dist は u と v の間の辺の数を返します。