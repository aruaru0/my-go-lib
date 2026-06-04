# MP — 文字列パターンマッチング (Morris-Pratt アルゴリズム)

## 型

- `MP` — パターンマッチャー構造体

## 関数

- `NewMP(pattern string) *MP` — 新しいパターンマッチャーを作成します
- `(*MP).FindAll(text string) []int` — テキスト内のパターンのすべての開始位置を返します

## 使用例

```go
mp := mylib.NewMP("abc")
pos := mp.FindAll("abcabcabc")
// pos = [0, 3, 6]
```
