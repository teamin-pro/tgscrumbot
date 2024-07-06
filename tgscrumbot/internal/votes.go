package internal

import "math"

// user - vote
type votes map[int64]int

func (v votes) add(user int64, vote int) {
	v[user] = vote
}

func (v votes) num() int {
	return len(v)
}

func (v votes) sum() int {
	sum := 0
	for _, vote := range v {
		sum += vote
	}
	return sum
}

func (v votes) avg() int {
	if len(v) == 0 {
		return 0
	}
	return int(math.Ceil(float64(v.sum()) / float64(v.num())))
}
