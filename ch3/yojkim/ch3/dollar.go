package ch3

type Dollar struct {
	Amount int
}

func NewDollar(amount int) Dollar {
	dollar := Dollar{}
	dollar.Amount = amount

	return dollar
}

func (d Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.Amount * multiplier}
}

func (d Dollar) Equal(input Dollar) bool {
	return d.Amount == input.Amount
}
