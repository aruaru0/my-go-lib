# LIS

最長増加部分列 (Longest Increasing Subsequence) を計算します。ジェネリクスに対応しています。

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
