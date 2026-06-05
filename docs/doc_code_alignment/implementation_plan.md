# 実装計画：ソースコードとドキュメントの不整合解消およびバグ修正

本計画は、先ほどの調査レポートで指摘したすべての課題を解消するための修正方針を示したものです。

## ユーザーレビューが必要な項目

> [!IMPORTANT]
> 1. **`gcdlcm.md` の削除と `gcd.md` への統合**:
>    ファイル名ベースの不一致を避けるため、不要な `gcdlcm.md` (EN/JP) は削除し、ソースコード `gcd.go` と同名の `gcd.md` (EN/JP) の中に可変引数 `Lcm` の説明を統合します。
> 2. **`Treap.Len()` の挙動変更**:
>    ドキュメントの定義「ユニークなキーの数を返す（returns the number of distinct keys）」に合わせるため、`Insert`/`Delete` 時に実際にキーの種類数（ユニーク数）が増減したタイミングでのみ内部サイズ（`t.size`）を増減させるように修正します。

---

## 提案する変更内容

### 1. ソースコードの修正

#### [MODIFY] [gcd.go](file:///Users/tadanori/opencode/my-go-lib/src/gcd.go)
- `Lcm` 関数を可変引数に対応するように拡張します。
```go
func Lcm[T Integer](a, b T, integers ...T) T {
	res := a / Gcd(a, b) * b
	for _, val := range integers {
		res = res / Gcd(res, val) * val
	}
	return res
}
```

#### [MODIFY] [gcd_test.go](file:///Users/tadanori/opencode/my-go-lib/src/gcd_test.go)
- `Lcm` 関数の可変引数を検証するテストを追加します（例: `Lcm(2, 3, 4) == 12`）。

#### [MODIFY] [treap.go](file:///Users/tadanori/opencode/my-go-lib/src/treap.go)
- `insert` および `delete` ヘルパー関数が、それぞれ「新規ノードが挿入されたか」「ノードが完全に削除（消滅）したか」の `bool` フラグを返すように修正します。
- `Insert` および `Delete` メソッドにおいて、この `bool` フラグをもとに `t.size` を正しく増減させることで、重複キー挿入や存在しないキーの削除で `t.size` が狂うバグを修正します。

#### [MODIFY] [treap_test.go](file:///Users/tadanori/opencode/my-go-lib/src/treap_test.go)
- 重複キー挿入時や存在しないキーの削除時における `Len()` の一貫性を検証するテストを追加します。

---

### 2. ドキュメントの修正および新規作成

#### [NEW] [types.md](file:///Users/tadanori/opencode/my-go-lib/doc/types.md) / [types.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/types.md)
- `src/types.go` で定義されているジェネリクス型制約 `Ordered` と `Number` についての説明ドキュメントを新規作成します。

#### [DELETE] [gcdlcm.md](file:///Users/tadanori/opencode/my-go-lib/doc/gcdlcm.md) / [gcdlcm.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/gcdlcm.md)
- 重複およびファイル名乖離を避けるため削除します。

#### [MODIFY] [gcd.md](file:///Users/tadanori/opencode/my-go-lib/doc/gcd.md) / [gcd.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/gcd.md)
- 可変引数 LCM のサポートに合わせてシグネチャおよび Example を更新します。

#### [MODIFY] [ahocorasick.md](file:///Users/tadanori/opencode/my-go-lib/doc/ahocorasick.md) / [ahocorasick.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/ahocorasick.md)
- 未記載メソッド `GetMask` の説明を追加します。

#### [MODIFY] [mincostflow.md](file:///Users/tadanori/opencode/my-go-lib/doc/mincostflow.md) / [mincostflow.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/mincostflow.md)
- 未記載メソッド `GetEdge` および `Edges` の説明を追加します。

#### [MODIFY] [priorityqueue.md](file:///Users/tadanori/opencode/my-go-lib/doc/priorityqueue.md) / [priorityqueue.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/priorityqueue.md)
- 未記載メソッド `Peek` の説明を追加します。

#### [MODIFY] [treap.md](file:///Users/tadanori/opencode/my-go-lib/doc/treap.md) / [treap.md (JP)](file:///Users/tadanori/opencode/my-go-lib/doc_jp/treap.md)
- 未記載メソッド `Kth` の説明を追加します。

---

## 検証計画

### 自動テスト
- `src` ディレクトリ内で `go test -v .` を実行し、すべてのテストがパスすることを確認します。
- 静的解析警告がないことを `go vet` で確認します。
- `check_signatures.go` スクリプトを再実行し、不一致の検出結果がゼロ（または `main` や簡易解析起因の誤検知のみ）になることを確認します。
