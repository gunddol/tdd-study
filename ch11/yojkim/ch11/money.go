package ch11

type Money struct {
	currency string
	amount int
}

func NewMoney(amount int, currency string) Money {
	return Money{
		amount: amount,
		currency: currency,
	}
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount * multiplier, m.currency)
}

func (m Money) Equal(input Money) bool {
	if m.currency != input.currency {
		return false
	}
	return m.amount == input.amount
}

func (m Money) Currency() string {
	return m.currency
}