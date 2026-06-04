# EulerPhi

オイラーのトーシェント関数 (φ 関数) です。ジェネリクスに対応しています。

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	fmt.Println(mylib.EulerPhi(12)) // 4
	fmt.Println(mylib.EulerPhi(7))  // 6
}
```

## 関数シグネチャ

```go
func EulerPhi[T Integer](n T) T
```

## アルゴリズム

- φ(n) = n × Π(1 - 1/p) を n の各素因数 p について計算します (O(√n))
