# Sliding Window (スライド最小値・最大値)

ジェネリクス対応のスライドウィンドウ最小値/最大値を O(n) で求める関数です。

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(mylib.SlideMin(a, 3)) // [-1 -3 -3 -3 3 3]
	fmt.Println(mylib.SlideMax(a, 3)) // [3 3 5 5 6 7]
}
```

## 関数シグネチャ

```go
func SlideMin[T Ordered](a []T, k int) []T
func SlideMax[T Ordered](a []T, k int) []T
```

## アルゴリズム

- 両端キュー (deque) を用いた O(n) のスライド最小値/最大値
- 蟻本 4.4 P.300 を元に実装
