# MultiSet — ジェネリックマルチセット

`map[T]int` を内部で使用したシンプルなマルチセットです。要素の追加と削除、カウントの追跡をサポートします。

## 関数

- `NewMultiSet[T comparable]() *MultiSet[T]` — 空のマルチセットを作成します
- `(*MultiSet[T]).Put(x T)` — `x` のインスタンスを 1 つ追加します
- `(*MultiSet[T]).Remove(x T)` — `x` のインスタンスを 1 つ削除します
- `(*MultiSet[T]).Count(x T) int` — `x` の個数を返します
- `(*MultiSet[T]).Values() []T` — すべての異なる要素を返します
- `(*MultiSet[T]).Len() int` — 異なる要素の数を返します
