# SuffixArray — SA-IS 接尾辞配列構築

文字列のすべての接尾辞（Suffix）を辞書順にソートしてインデックスを記録した配列で、任意のパターン検索を高速に行うためのデータ構造です。

## 主なユースケース（どういう時に使うか）

- **巨大なテキスト（ゲノムデータや書籍データなど）から、任意のキーワードを何度も高速に全文検索する処理**
- 文字列中の最長反復部分文字列の探索

## なぜ使うのか（メリット）

- 接尾辞配列を構築（$O(N \log N)$ または $O(N)$）しておくことで、任意のパターン $P$ の検索を $O(|P| \log N)$ の二分探索で実行できます。インデックス作成後は省メモリで動作する点が接尾辞木に比べて有利です。

---

## 関数

- `SuffixArray(s []int, upper int) []int` — [0, upper] の値を持つ整数スライスの接尾辞配列を返します
- `SuffixArrayString(s string) []int` — 文字列 (バイト値) の接尾辞配列を返します
- `SuffixArrayInt(s []int) []int` — 任意の整数スライスの接尾辞配列を返します (内部で値の圧縮を行います)
- `LcpArray(s, sa []int) []int` — 整数スライスとその接尾辞配列から LCP 配列を計算します
- `LcpArrayString(s string, sa []int) []int` — 文字列とその接尾辞配列から LCP 配列を計算します

## 使用例

```go
sa := mylib.SuffixArrayString("banana")
// sa = [5 3 1 0 4 2]

lcp := mylib.LcpArrayString("banana", sa)
// lcp = [1 3 0 0 2]
```