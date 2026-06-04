# my-go-lib
競技プログラミングのためのGoライブラリ集。AtCoder で使われる典型的なアルゴリズム・データ構造を提供します。

これまでテンプレートとして保存していた、Go言語のアルゴリズム・データ構造をライブラリにまとめたものです。


## インストール

```bash
go get github.com/aruaru0/my-go-lib
```

## 使い方

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	// GCD / LCM
	fmt.Println(mylib.Gcd(12, 8))  // 4
	fmt.Println(mylib.Lcm(12, 8)) // 24

	// 素因数分解
	fmt.Println(mylib.PrimeFactorMap(12)) // map[2:2 3:1]

	// UnionFind (DSU)
	uf := mylib.NewUnionFind(5)
	uf.Unite(0, 1)
	fmt.Println(uf.Same(0, 1)) // true

	// PriorityQueue
	pq := mylib.NewPriorityQueue[int](func(a, b int) bool { return a < b })
	pq.Push(3)
	pq.Push(1)
	pq.Push(2)
	fmt.Println(pq.Pop()) // 1

	// セグメント木
	e := func() int { return 0 }
	merger := func(a, b int) int { return a + b }
	seg := mylib.NewSegTree([]int{1, 2, 3, 4, 5}, e, merger)
	fmt.Println(seg.Prod(1, 4)) // 9 (2+3+4)
}
```

## モジュール一覧

### 数学

| モジュール       | 説明                                       | Docs                                                   |
| ---------------- | ------------------------------------------ | ------------------------------------------------------ |
| GCD / LCM        | 最大公約数・最小公倍数（ジェネリクス対応） | [EN](doc/gcd.md) / [JP](doc_jp/gcd.md)                 |
| 素因数分解       | 試し割りによる素因数分解                   | [EN](doc/primefactor.md) / [JP](doc_jp/primefactor.md) |
| 約数列挙         | 約数の列挙                                 | [EN](doc/divisor.md) / [JP](doc_jp/divisor.md)         |
| オイラーのφ関数  | トーシェント関数                           | [EN](doc/eulerphi.md) / [JP](doc_jp/eulerphi.md)       |
| nCk / nPk / nHk  | 階乗テーブルによる組み合わせ計算 (mod)     | [EN](doc/nck.md) / [JP](doc_jp/nck.md)                 |
| nCk テーブル     | パスカルの三角形                           | [EN](doc/combtable.md) / [JP](doc_jp/combtable.md)     |
| nCk自動生成      | 自動拡張する階乗テーブル                   | [EN](doc/ncrautogen.md) / [JP](doc_jp/ncrautogen.md)   |
| modint           | mod演算 (加減乗除・累乗・行列累乗)         | [EN](doc/modint.md) / [JP](doc_jp/modint.md)           |
| 拡張ユークリッド | 拡張ユークリッド互除法                     | [EN](doc/extgcd.md) / [JP](doc_jp/extgcd.md)           |
| 逆元列挙         | 1..n-1 の逆元を O(n) で前計算              | [EN](doc/invmod.md) / [JP](doc_jp/invmod.md)           |
| 中国剰余定理     | Garner のアルゴリズム                      | [EN](doc/garner.md) / [JP](doc_jp/garner.md)           |
| FFT              | 高速フーリエ変換 (complex128)              | [EN](doc/fft.md) / [JP](doc_jp/fft.md)                 |
| 畳み込み         | NTT による畳み込み                         | [EN](doc/convolution.md) / [JP](doc_jp/convolution.md) |
| 行列累乗         | 行列の累乗 (ジェネリクス)                  | [EN](doc/matrixpow.md) / [JP](doc_jp/matrixpow.md)     |
| 3x3行列積        | 3x3 行列の積                               | [EN](doc/matrix3.md) / [JP](doc_jp/matrix3.md)         |
| FloorSum         | floor sum / 各種数論関数                   | [EN](doc/mathutil.md) / [JP](doc_jp/mathutil.md)       |
| XOR総和          | 全ペアの XOR 和                            | [EN](doc/xorsum.md) / [JP](doc_jp/xorsum.md)           |
| bits             | ceilPow2, bsf                              | [EN](doc/bits.md) / [JP](doc_jp/bits.md)               |

### データ構造

| モジュール      | 説明                              | Docs                                                       |
| --------------- | --------------------------------- | ---------------------------------------------------------- |
| UnionFind (DSU) | 素集合データ構造                  | [EN](doc/unionfind.md) / [JP](doc_jp/unionfind.md)         |
| DSU (ACL版)     | Groups() 付き UnionFind           | [EN](doc/dsu.md) / [JP](doc_jp/dsu.md)                     |
| BIT (Fenwick)   | Binary Indexed Tree (int)         | [EN](doc/bit.md) / [JP](doc_jp/bit.md)                     |
| FenwickTree     | ジェネリクス対応 BIT + LowerBound | [EN](doc/fenwick.md) / [JP](doc_jp/fenwick.md)             |
| SegTree         | セグメント木 (ジェネリクス)       | [EN](doc/segtree.md) / [JP](doc_jp/segtree.md)             |
| LazySegTree     | 遅延評価セグメント木              | [EN](doc/lazysegtree.md) / [JP](doc_jp/lazysegtree.md)     |
| PriorityQueue   | 二分ヒープ (ジェネリクス)         | [EN](doc/priorityqueue.md) / [JP](doc_jp/priorityqueue.md) |
| Treap           | 平衡二分探索木 ( randomized BST ) | [EN](doc/treap.md) / [JP](doc_jp/treap.md)                 |
| Set             | 集合 (ジェネリクス)               | [EN](doc/set.md) / [JP](doc_jp/set.md)                     |
| MultiSet        | 多重集合                          | [EN](doc/multiset.md) / [JP](doc_jp/multiset.md)           |
| OrderedMap      | 挿入順序保持マップ                | [EN](doc/orderedmap.md) / [JP](doc_jp/orderedmap.md)       |
| Median          | 動的中央値                        | [EN](doc/median.md) / [JP](doc_jp/median.md)               |

### グラフ

| モジュール       | 説明                            | Docs                                                             |
| ---------------- | ------------------------------- | ---------------------------------------------------------------- |
| BellmanFord      | ベルマンフォード法 (負閉路検出) | [EN](doc/bellmanford.md) / [JP](doc_jp/bellmanford.md)           |
| BellmanFordRoute | 経路付きベルマンフォード        | [EN](doc/bellmanfordroute.md) / [JP](doc_jp/bellmanfordroute.md) |
| DFS Route        | DFS による経路探索              | [EN](doc/dfsroute.md) / [JP](doc_jp/dfsroute.md)                 |
| SCC              | 強連結成分分解                  | [EN](doc/scc.md) / [JP](doc_jp/scc.md)                           |
| TwoSAT           | 2-SAT ソルバー                  | [EN](doc/twosat.md) / [JP](doc_jp/twosat.md)                     |
| MaxFlow (Dinic)  | 最大流                          | [EN](doc/maxflow.md) / [JP](doc_jp/maxflow.md)                   |
| Dinic            | Dinic (シンプル版)              | [EN](doc/dinic.md) / [JP](doc_jp/dinic.md)                       |
| MinCostFlow      | 最小費用流                      | [EN](doc/mincostflow.md) / [JP](doc_jp/mincostflow.md)           |
| LCA              | 最小共通祖先 (ダブリング)       | [EN](doc/lca.md) / [JP](doc_jp/lca.md)                           |
| LCASeg           | LCA (セグメント木 + Euler Tour) | [EN](doc/lcaseg.md) / [JP](doc_jp/lcaseg.md)                     |
| 全方位木DP       | 全方位木 DP (木の距離)          | [EN](doc/rerootdp.md) / [JP](doc_jp/rerootdp.md)                 |
| ConvexHullTrick  | Li Chao セグメント木            | [EN](doc/convexhulltrick.md) / [JP](doc_jp/convexhulltrick.md)   |
| KthStep          | ダブリングによる K ステップ移動 | [EN](doc/kthstep.md) / [JP](doc_jp/kthstep.md)                   |
| KthStepLoop      | ループ検出付き K ステップ       | [EN](doc/kthsteploop.md) / [JP](doc_jp/kthsteploop.md)           |
| MaxGain          | K 回移動の最大利得              | [EN](doc/maxgain.md) / [JP](doc_jp/maxgain.md)                   |
| Clique           | クリーク被覆数 (bit DP)         | [EN](doc/clique.md) / [JP](doc_jp/clique.md)                     |

### 文字列

| モジュール  | 説明                          | Docs                                                   |
| ----------- | ----------------------------- | ------------------------------------------------------ |
| MP          | Morris-Pratt 文字列マッチング | [EN](doc/mp.md) / [JP](doc_jp/mp.md)                   |
| KMP         | KMP 法 (ジェネリクス)         | [EN](doc/kmp.md) / [JP](doc_jp/kmp.md)                 |
| AhoCorasick | 複数パターンマッチング        | [EN](doc/ahocorasick.md) / [JP](doc_jp/ahocorasick.md) |
| RollingHash | ローリングハッシュ            | [EN](doc/rollinghash.md) / [JP](doc_jp/rollinghash.md) |
| ZAlgorithm  | Z アルゴリズム                | [EN](doc/zalgorithm.md) / [JP](doc_jp/zalgorithm.md)   |
| SuffixArray | SA-IS 接尾辞配列構築          | [EN](doc/suffixarray.md) / [JP](doc_jp/suffixarray.md) |
| Levenshtein | 編集距離                      | [EN](doc/levenshtein.md) / [JP](doc_jp/levenshtein.md) |

### DP・探索

| モジュール     | 説明                      | Docs                                                           |
| -------------- | ------------------------- | -------------------------------------------------------------- |
| LIS            | 最長増加部分列            | [EN](doc/lis.md) / [JP](doc_jp/lis.md)                         |
| 桁DP           | 数値DP テンプレート       | [EN](doc/digitdp.md) / [JP](doc_jp/digitdp.md)                 |
| スライド最小値 | スライド窓の最小/最大値   | [EN](doc/slidingwindow.md) / [JP](doc_jp/slidingwindow.md)     |
| 順列列挙       | next_permutation          | [EN](doc/nextpermutation.md) / [JP](doc_jp/nextpermutation.md) |
| 組み合わせ列挙 | nCr 列挙 (チャネル)       | [EN](doc/combinations.md) / [JP](doc_jp/combinations.md)       |
| 部分集合列挙   | ビット部分集合の列挙      | [EN](doc/subsetenum.md) / [JP](doc_jp/subsetenum.md)           |
| 数列巡回       | MOD 巡回パターンの総和    | [EN](doc/cyclesum.md) / [JP](doc_jp/cyclesum.md)               |
| N進数列挙      | N進数 n 桁列挙            | [EN](doc/nary.md) / [JP](doc_jp/nary.md)                       |
| 座標圧縮       | 値のマッピング (座標圧縮) | [EN](doc/compress.md) / [JP](doc_jp/compress.md)               |
| 区間分割       | P で割り切れる部分列      | [EN](doc/divsubstr.md) / [JP](doc_jp/divsubstr.md)             |

### 幾何

| モジュール | 説明           | Docs                                                 |
| ---------- | -------------- | ---------------------------------------------------- |
| 凸包       | Monotone Chain | [EN](doc/convexhull.md) / [JP](doc_jp/convexhull.md) |
| 円の交点   | 2 つの円の交点 | [EN](doc/circle.md) / [JP](doc_jp/circle.md)         |
| 線分の交差 | 線分交差判定   | [EN](doc/segment.md) / [JP](doc_jp/segment.md)       |

### その他

| モジュール | 説明                   | Docs                                             |
| ---------- | ---------------------- | ------------------------------------------------ |
| Hash       | 簡易ハッシュ関数       | [EN](doc/hash.md) / [JP](doc_jp/hash.md)         |
| BaseConv   | n進数変換 (負数対応)   | [EN](doc/baseconv.md) / [JP](doc_jp/baseconv.md) |
| Gacha      | ガチャコンプ期待値     | [EN](doc/gacha.md) / [JP](doc_jp/gacha.md)       |
| Median     | 動的中央値 (追加/削除) | [EN](doc/median.md) / [JP](doc_jp/median.md)     |

## ドキュメント

- [English](doc/) — 英語ドキュメント
- [日本語](doc_jp/) — 日本語ドキュメント

## テスト

```bash
cd src && go test ./...
```
