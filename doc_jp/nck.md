# NCR

事前計算された階乗を使用して、nCk、nPk、nHk を指定された法で計算します。法はコンストラクタで指定します。

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	nc := mylib.NewNCR(1000000007, 100)
	fmt.Println(nc.N(5, 2))  // 10 (nCk)
	fmt.Println(nc.P(5, 2))  // 20 (nPk)
	fmt.Println(nc.H(5, 2))  // 6  (nHk)
}
```

## 関数シグネチャ

```go
func NewNCR(mod int, maxN int) *NCR
func (nc *NCR) N(n, k int) int
func (nc *NCR) P(n, k int) int
func (nc *NCR) H(n, k int) int
```

## アルゴリズム

- 階乗と逆階乗を O(maxN) で事前計算
- フェルマーの小定理によるモジュラ逆数
- 各クエリは O(1)
