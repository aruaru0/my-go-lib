# NextPermutation

与えられた配列の要素の並び順について、辞書順で「直後（次）」となる順列（Permutation）に配列をインプレースで並び替えるアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- **$N$ 個の異なる要素を並べ替えるすべてのパターン（$N!$ 通り）を順次生成して全探索する処理**
- 巡回セールスマン問題などで小規模な経路の総当たり探索

## なぜ使うのか（メリット）

- 再帰を使わず、現在の配列のみから $O(N)$ の時間計算量かつ $O(1)$ の追加空間で次の順列を順次生成できるため、メモリ効率が極めて高いです。

---

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