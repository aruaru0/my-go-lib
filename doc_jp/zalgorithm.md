# ZAlgorithm — Z アルゴリズム (ジェネリック)

文字列 $S$ に対し、 $S$ の各接頭辞と、 $S$ の各開始位置から始まる部分文字列との最長共通接頭辞（LCP）の長さを求めるアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- **文字列探索（検索パターン $P$ とテキスト $T$ を連結して Z-algorithm を適用）**
- 文字列内の最長繰り返しパターンや接頭辞の周期の検出

## なぜ使うのか（メリット）

- 文字列の長さを $N$ としたとき、一度の走査で $O(N)$ という線形時間ですべての位置における最長共通接頭辞長を漏れなく計算できます。

---

## 関数

- `ZAlgorithm[T comparable](s []T) []int` — comparable なスライスの Z 配列を計算します
- `ZAlgorithmString(s string) []int` — 文字列の Z 配列を計算します

Z 配列 `z` において、`z[i]` は `s` と `s[i:]` の最長共通接頭辞の長さです。

## 使用例

```go
z := mylib.ZAlgorithmString("aaaaa")
// z = [5 4 3 2 1]

z2 := mylib.ZAlgorithmString("abcabcabc")
// z2 = [9 0 0 6 0 0 3 0 0]
```