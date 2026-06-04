# CombTable

パスカルの三角形を使用して計算された二項係数 nCk のテーブルです。

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
