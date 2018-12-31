package ch8

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount}
}

func (d Dollar) Times(multiplier int) Money {
	return Dollar{d.amount * multiplier}
}

func (d Dollar) Equal(input Money) bool {
	dollar, ok := input.(Dollar)
	if !ok {
		return false
	}
	return d.amount == dollar.amount
}
