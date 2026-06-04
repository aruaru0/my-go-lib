# LevenshteinDistance — 編集距離

## 関数

- `LevenshteinDistance(s, t string) int` — 2 つの文字列間の Levenshtein 距離 (1 文字編集の最小回数) を返します

## 使用例

```go
d := mylib.LevenshteinDistance("kitten", "sitting")
// d = 3
```
