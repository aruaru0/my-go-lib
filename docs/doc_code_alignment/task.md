# タスクリスト：ソースコードとドキュメントの不整合解消およびバグ修正

- [x] `src/gcd.go` の修正（Lcmを可変引数に）
- [x] `src/gcd_test.go` にテストコードを追加
- [x] `src/treap.go` の `Insert`/`Delete` における `size` 管理バグの修正
- [x] `src/treap_test.go` に要素数とバグの挙動検証用テストを追加
- [x] ドキュメントファイルの整理
  - [x] `doc/types.md` / `doc_jp/types.md` を新規作成
  - [x] `doc/gcdlcm.md` / `doc_jp/gcdlcm.md` を削除
  - [x] `doc/gcd.md` / `doc_jp/gcd.md` を更新
  - [x] 各種未記載メソッドの説明を追記
    - [x] `doc/ahocorasick.md` / `doc_jp/ahocorasick.md` (`GetMask`)
    - [x] `doc/mincostflow.md` / `doc_jp/mincostflow.md` (`GetEdge`, `Edges`)
    - [x] `doc/priorityqueue.md` / `doc_jp/priorityqueue.md` (`Peek`)
    - [x] `doc/treap.md` / `doc_jp/treap.md` (`Kth`)
- [x] 検証
  - [x] `go test` によるテスト全体の実行とパスの確認
  - [x] `go vet` による静的解析の確認
  - [x] `check_signatures.go` スクリプトの実行による不一致解消の確認
