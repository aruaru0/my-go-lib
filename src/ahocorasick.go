package mylib

type AhoMatch struct {
	Pos int
	ID  int
}

type AhoCorasick struct {
	to      []map[rune]int
	cnt     []int
	mask    []int
	fail    []int
	nextID  int
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		to:   []map[rune]int{{}},
		cnt:  []int{0},
		mask: []int{0},
	}
}

func (ac *AhoCorasick) Add(pattern string) int {
	id := ac.nextID
	ac.nextID++
	v := 0
	for _, c := range pattern {
		if _, ok := ac.to[v][c]; !ok {
			ac.to[v][c] = len(ac.to)
			ac.to = append(ac.to, map[rune]int{})
			ac.cnt = append(ac.cnt, 0)
			ac.mask = append(ac.mask, 0)
		}
		v = ac.to[v][c]
	}
	ac.cnt[v]++
	ac.mask[v] |= (1 << id)
	return id
}

func (ac *AhoCorasick) Build() {
	ac.fail = make([]int, len(ac.to))
	for i := range ac.fail {
		ac.fail[i] = -1
	}
	q := []int{0}
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		for c, u := range ac.to[v] {
			ac.fail[u] = ac.nextState(ac.fail[v], c)
			ac.cnt[u] += ac.cnt[ac.fail[u]]
			ac.mask[u] |= ac.mask[ac.fail[u]]
			q = append(q, u)
		}
	}
}

func (ac *AhoCorasick) nextState(v int, c rune) int {
	for v != -1 {
		if u, ok := ac.to[v][c]; ok {
			return u
		}
		v = ac.fail[v]
	}
	return 0
}

func (ac *AhoCorasick) Search(text string) []AhoMatch {
	var matches []AhoMatch
	v := 0
	for i, c := range text {
		v = ac.nextState(v, c)
		if m := ac.mask[v]; m != 0 {
			for id := 0; id < ac.nextID; id++ {
				if m>>id&1 != 0 {
					matches = append(matches, AhoMatch{Pos: i, ID: id})
				}
			}
		}
	}
	return matches
}

func (ac *AhoCorasick) MatchCount(text string) int {
	v := 0
	total := 0
	for _, c := range text {
		v = ac.nextState(v, c)
		total += ac.cnt[v]
	}
	return total
}

func (ac *AhoCorasick) GetMask(v int) int {
	return ac.mask[v]
}
