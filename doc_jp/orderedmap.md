# OrderedMap — 挿入順序を保持するマップ

スライスとマップの組み合わせを使用して、キーと値のペアを挿入順に管理します。

## 関数

- `NewOrderedMap[K comparable, V any]() *OrderedMap[K, V]` — 空の順序付きマップを作成します
- `(*OrderedMap[K, V]).Set(key K, value V)` — キーを設定します (元の挿入順序を保持します)
- `(*OrderedMap[K, V]).Get(key K) (V, bool)` — キーで値を取得します
- `(*OrderedMap[K, V]).Delete(key K)` — キーを削除します
- `(*OrderedMap[K, V]).Keys() []K` — キーを挿入順で返します
- `(*OrderedMap[K, V]).Len() int` — エントリ数を返します
