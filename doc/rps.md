# RPS — Rock-Paper-Scissors

## Functions

- `RPSWin[T ~byte](a, b T) T` — returns the winning hand given two RPS moves. Rules: R > S, S > P, P > R. If tie, returns the same move. Accepts any type with underlying type `byte`.
