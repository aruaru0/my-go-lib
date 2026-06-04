# AhoCorasick — 複数パターン文字列マッチング

## 型

- `AhoCorasick` — Aho-Corasick オートマトン
- `AhoMatch` — 位置におけるマッチしたパターンを表します

## 関数

- `NewAhoCorasick() *AhoCorasick` — 新しいオートマトンを作成します
- `(*AhoCorasick).Add(pattern string) int` — パターンを追加し、その ID を返します
- `(*AhoCorasick).Build()` — 失敗リンクを構築します (すべてのパターンを追加した後に呼び出します)
- `(*AhoCorasick).Search(text string) []AhoMatch` — すべてのマッチ (位置, ID) を返します
- `(*AhoCorasick).MatchCount(text string) int` — マッチの総数を返します

## 使用例

```go
ac := mylib.NewAhoCorasick()
ac.Add("he")
ac.Add("she")
ac.Add("hers")
ac.Build()
matches := ac.Search("ushers")
// matches: 位置 1 (she), 位置 2 (he), 位置 2 (hers)
```
