package ch15

type Bank struct {
	rates map[Pair]int
}

func NewBank() Bank {
	return Bank{
		rates: make(map[Pair]int),
	}
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[NewPair(from, to)]
}

func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[NewPair(from, to)] = rate
}