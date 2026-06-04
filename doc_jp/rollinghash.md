# RollingHash — ローリングハッシュ (Rabin–Karp)

## 関数

- `NewRollingHash(s string, base, mod int) *RollingHash` — 文字列 `s` に対して、カスタムの基数と法を使用したローリングハッシュを作成します。
- `NewRollingHashDefault(s string) *RollingHash` — 基数 37、法 1e9+7 でローリングハッシュを作成します。
- `(rh *RollingHash) Get(l, r int) int` — 部分文字列 `s[l:r)` のハッシュ値を返します。
- `(rh *RollingHash) Len() int` — 元の文字列の長さを返します。
