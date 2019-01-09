package ch12

type Money struct {
	currency string
	amount int
}

func newMoney(amount int, currency string) Money {
	return Money{
		amount: amount,
		currency: currency,
	}
}

func NewDollar(amount int) Money {
	return newMoney(amount, "USD")
}

func NewFranc(amount int) Money {
	return newMoney(amount, "CHF")
}

func (m Money) Times(multiplier int) Money {
	return newMoney(m.amount * multiplier, m.currency)
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

func (m Money) Plus(addend Money) Expression {
	return newMoney(m.amount + addend.amount, addend.currency)
}