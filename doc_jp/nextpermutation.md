# NextPermutation

次の辞書順順列を生成します。ジェネリクスに対応しています。

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	a := []int{1, 2, 3}
	for {
		fmt.Println(a)
		if !mylib.NextPermutation(a) {
			break
		}
	}
}
```

## 関数シグネチャ

```go
func NextPermutation[T Ordered](x []T) bool
```

`Ordered` はすべての comparable な基本型をサポートします。

## アルゴリズム

- C++ の `std::next_permutation` と同じアルゴリズム (O(n))
- 次の順列が存在する場合は `true`、そうでない場合は `false` を返します
