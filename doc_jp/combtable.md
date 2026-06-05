# CombTable

パスカルの三角形の原理を利用して、二項係数 $nCr$ （組み合わせの数）の値を事前にテーブルに計算しておくデータ構造です。

## 主なユースケース（どういう時に使うか）

- 競技プログラミングや動的計画法の中で、何度も異なる $n$ と $r$ について組み合わせ数 $nCr$ を高速に参照したい場合
- **剰余（Mod）を取らない、比較的小さな $n$ に対する組み合わせ数を求める場合**

## なぜ使うのか（メリット）

- 事前に $O(N^2)$ の前処理を行うことで、任意の $nCr$ のクエリに対して $O(1)$ の定数時間で結果を返せます。
- 整数演算のみで構築するため、浮動小数点数の誤差がありません。

---

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	ct := mylib.NewCombTable[int](10)
	fmt.Println(ct.NCk(5, 2))  // 10
	fmt.Println(ct.NCk(10, 5)) // 252
}
```

## 関数シグネチャ

```go
func NewCombTable[T Integer](n int) *CombTable[T]
func (c *CombTable[T]) NCk(n, k int) T
```

## アルゴリズム

- パスカルの三角形を使用して O(n²) で事前計算
- テーブルサイズ: (n+1) × (n+1)
- k < 0、n < k、または n がテーブルサイズを超える場合は 0 を返します