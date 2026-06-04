package mylib

func RPSWin[T ~byte](a, b T) T {
	if a == b {
		return a
	}
	switch {
	case a == 'R' && b == 'S':
		return a
	case a == 'R' && b == 'P':
		return b
	case a == 'S' && b == 'P':
		return a
	case a == 'S' && b == 'R':
		return b
	case a == 'P' && b == 'S':
		return b
	case a == 'P' && b == 'R':
		return a
	}
	return a
}
