# Divisors

整数のすべての約数を列挙します。ジェネリクスに対応しています。

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	divs := mylib.Divisors(12)
	fmt.Println(divs) // [1 12 2 6 3 4] (ソートされていません)
}
```

## 関数シグネチャ

```go
func Divisors[T Integer](n T) []T
```

## アルゴリズム

- i を 1 から √n までループし、n % i == 0 のときに i と n/i を収集します (O(√n))
- 結果はソートされていません
