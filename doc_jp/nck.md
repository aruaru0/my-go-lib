# NCR

事前に階乗とその逆元を計算（前処理）しておくことで、任意の $N$ と $K$ に対する二項係数 $nCr \pmod P$ を高速に計算するライブラリです。

## 主なユースケース（どういう時に使うか）

- **競技プログラミングにおいて、巨大な $N, K$ （例: $10^6$ 程度まで）に対する組み合わせ数を何度もクエリで求める場合**

## なぜ使うのか（メリット）

- $O(N)$ の時間で階乗およびその逆元のテーブルを前処理しておくことにより、クエリごとの組み合わせ計算を $O(1)$ という定数時間で終了させられます。

---

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