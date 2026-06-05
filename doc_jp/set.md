# Set — ジェネリック集合

重複する要素を持たない値の集まり（Set）を管理し、要素の追加・削除・存在判定、および集合同士の和集合・積集合などの基本操作をサポートするデータ構造です。

## 主なユースケース（どういう時に使うか）

- 重複データを排除したユニークなIDや項目の管理
- **2つのユーザーリストにおける共通のフォロワー（積集合）や全メンバー（和集合）の抽出**

## なぜ使うのか（メリット）

- ハッシュマップをベースに構築され、各要素の追加、削除、存在判定を平均 $O(1)$ という極めて高いパフォーマンスで行えます。

---

## 関数

- `NewSet[T comparable]() Set[T]` — 空の集合を作成します
- `NewSetFromSlice[T comparable](s []T) Set[T]` — スライスから集合を作成します
- `(Set[T]).Add(v T) bool` — 値を追加します。新しく追加された場合は true を返します
- `(Set[T]).Remove(v T)` — 値を削除します
- `(Set[T]).Contains(v T) bool` — 値の有無を確認します
- `(Set[T]).Cardinality() int` — 要素数を返します
- `(Set[T]).Clear()` — すべての要素を削除します
- `(Set[T]).Values() []T` — すべての要素をスライスとして返します
- `(Set[T]).Union(other Set[T]) Set[T]` — 新しい和集合を返します
- `(Set[T]).Intersect(other Set[T]) Set[T]` — 新しい積集合を返します
- `(Set[T]).Difference(other Set[T]) Set[T]` — 新しい差集合を返します
- `(Set[T]).Equal(other Set[T]) bool` — 集合の等価性を確認します