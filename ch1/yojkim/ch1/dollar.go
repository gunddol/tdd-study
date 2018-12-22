package ch1

type Dollar struct {
	Amount int
}

func NewDollar(amount int) Dollar {
	dollar := Dollar{}
	dollar.Amount = 10

	return dollar
}

func (d Dollar) Times(multiplier int) {}
