# RollingHash — Rolling Hash (Rabin–Karp)

## Functions

- `NewRollingHash(s string, base, mod int) *RollingHash` — creates a rolling hash for string `s` with custom base and mod.
- `NewRollingHashDefault(s string) *RollingHash` — creates a rolling hash with base=37 and mod=1e9+7.
- `(rh *RollingHash) Get(l, r int) int` — returns the hash of substring `s[l:r)`.
- `(rh *RollingHash) Len() int` — returns the length of the original string.
