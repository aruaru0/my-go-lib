# LIS

与えられた数列の中から、要素の順序を崩さずに（ただし連続していなくてもよい）要素の値が狭義（または広義）に増加していく部分列のうち、最も長いものの長さを求めるアルゴリズムです（最長増加部分列）。

## 主なユースケース（どういう時に使うか）

- データの並びから増加傾向にある最長のパターンを特定するトレンド分析
- カードや荷物の並び替えを最小の手順で行うタスクの解決
- 2次元平面上での干渉しない区間の最大数選定

## なぜ使うのか（メリット）

- 二分探索（`sort.Search`など）を組み合わせた動的計画法により、全体で $O(N \log N)$ という非常に高速な計算量で最長長さを求められます。

---

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(mylib.Lis(a)) // 4

	// LowerBound も利用可能です
	b := []int{1, 3, 5, 7, 9}
	fmt.Println(mylib.LowerBound(b, 5)) // 2
}
```

## 関数シグネチャ

```go
func LowerBound[T Ordered](a []T, x T) int
func Lis[T Ordered](a []T) int
```

## アルゴリズム

- 二分探索を用いた O(n log n) の LIS (狭義単調増加)
- `LowerBound` はソートされたスライス内で値が x 以上となる最初の位置を返します