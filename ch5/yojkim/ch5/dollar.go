package ch5

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	dollar := Dollar{}
	dollar.amount = amount

	return dollar
}

func (d Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}

func (d Dollar) Equal(input Dollar) bool {
	return d.amount == input.amount
}
