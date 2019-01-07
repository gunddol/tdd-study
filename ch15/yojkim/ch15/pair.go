package ch15

type Pair struct {
	from string
	to string
}

func NewPair(from, to string) Pair {
	return Pair{
		from: from,
		to: to,
	}
}