package ch9

type Dollar struct {
	amount int
	currency string
}

func NewDollar(amount int) Dollar {
	return Dollar{
		amount: amount,
		currency: "USD",
	}
}

func (d Dollar) Times(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}

func (d Dollar) Equal(input Money) bool {
	dollar, ok := input.(Dollar)
	if !ok {
		return false
	}
	return d.amount == dollar.amount
}

func (d Dollar) Currency() string {
	return d.currency
}